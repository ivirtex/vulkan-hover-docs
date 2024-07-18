# VK_NV_acquire_winrt_display(3) Manual Page

## Name

VK_NV_acquire_winrt_display - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

346

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_direct_mode_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_direct_mode_display.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Juliano <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_acquire_winrt_display%5D%20@jjuliano%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_acquire_winrt_display%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jjuliano</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-09-29

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Juliano, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to take exclusive control of a
display on Windows 10 provided that the display is not already
controlled by a compositor. Examples of compositors include the Windows
desktop compositor, other applications using this Vulkan extension, and
applications that
[“Acquire”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaymanager.tryacquiretarget)
a
[“DisplayTarget”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget)
using a [“WinRT”](https://docs.microsoft.com/en-us/uwp/api/) command
such as
[“winrt::Windows::Devices::Display::Core::DisplayManager.TryAcquireTarget()”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaymanager.tryacquiretarget).

When control is acquired the application has exclusive access to the
display until control is released or the application terminates. An
application’s attempt to acquire is denied if a different application
has already acquired the display.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireWinrtDisplayNV.html)

- [vkGetWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetWinrtDisplayNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_ACQUIRE_WINRT_DISPLAY_EXTENSION_NAME`

- `VK_NV_ACQUIRE_WINRT_DISPLAY_SPEC_VERSION`

## <a href="#_issues" class="anchor"></a>Issues

1\) What should the platform substring be for this extension:

**RESOLVED**: The platform substring is “Winrt”.

The substring “Winrt” matches the fact that the OS API exposing the
acquire and release functionality is called “WinRT”.

The substring “Win32” is wrong because the related “WinRT” API is
explicitly **not** a “Win32” API. “WinRT” is a competing API family to
the “Win32” API family.

The substring “Windows” is suboptimal because there could be more than
one relevant API on the Windows platform. There is preference to use the
more-specific substring “Winrt”.

2\) Should [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireWinrtDisplayNV.html) take
a winRT DisplayTarget, or a Vulkan display handle as input?

**RESOLVED**: A Vulkan display handle. This matches the design of
[vkAcquireXlibDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireXlibDisplayEXT.html).

3\) Should the acquire command be platform-independent named
“vkAcquireDisplayNV”, or platform-specific named
“vkAcquireWinrtDisplayNV”?

**RESOLVED**: Add a platform-specific command.

The inputs to the Acquire command are all Vulkan types. None are WinRT
types. This opens the possibility of the winrt extension defining a
platform-independent acquire command.

The X11 acquire command does need to accept a platform-specific
parameter. This could be handled by adding to a platform-independent
acquire command a params struct to which platform-dependent types can be
chained by `pNext` pointer.

The prevailing opinion is that it would be odd to create a second
platform-independent function that is used on the Windows 10 platform,
but that is not used for the X11 platform. Since a Windows 10
platform-specific command is needed anyway for converting between
vkDisplayKHR and platform-native handles, opinion was to create a
platform-specific acquire function.

4\) Should the [vkGetWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetWinrtDisplayNV.html) parameter
identifying a display be named “deviceRelativeId” or
“adapterRelativeId”?

**RESOLVED**: The WinRT name is “AdapterRelativeId”. The name “adapter”
is the Windows analog to a Vulkan “physical device”. Vulkan already has
precedent to use the name `deviceLUID` for the concept that Windows APIs
call “AdapterLuid”. Keeping form with this precedent, the name
“deviceRelativeId” is chosen.

5\) Does [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireWinrtDisplayNV.html) cause
the Windows desktop compositor to release a display?

**RESOLVED**: No.
[vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireWinrtDisplayNV.html) does not itself
cause the Windows desktop compositor to release a display. This action
must be performed outside of Vulkan.

Beginning with Windows 10 version 2004 it is possible to cause the
Windows desktop compositor to release a display by using the “Advanced
display settings” sub-page of the “Display settings” control panel. See
<a
href="https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors"
class="bare">https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors</a>

6\) Where can one find additional information about custom compositors
for Windows 10?

**RESOLVED**: Relevant references are as follows.

According to Microsoft’s documentation on ["building a custom
compositor"](https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors-compositor),
the ability to write a custom compositor is not a replacement for a
fullscreen desktop window. The feature is for writing compositor apps
that drive specialized hardware.

Only certain editions of Windows 10 support custom compositors,
["documented
here"](https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors#windows-10-version-2004).
The product type can be queried from Windows 10. See <a
href="https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo"
class="bare">https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo</a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-09-29 (Jeff Juliano)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_acquire_winrt_display"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
