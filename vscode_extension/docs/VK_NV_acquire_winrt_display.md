# VK\_NV\_acquire\_winrt\_display(3) Manual Page

## Name

VK\_NV\_acquire\_winrt\_display - device extension



## [](#_registered_extension_number)Registered Extension Number

346

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_direct\_mode\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_direct_mode_display.html)

## [](#_contact)Contact

- Jeff Juliano [\[GitHub\]jjuliano](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_acquire_winrt_display%5D%20%40jjuliano%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_acquire_winrt_display%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-09-29

**IP Status**

No known IP claims.

**Contributors**

- Jeff Juliano, NVIDIA

## [](#_description)Description

This extension allows an application to take exclusive control of a display on Windows 10 provided that the display is not already controlled by a compositor. Examples of compositors include the Windows desktop compositor, other applications using this Vulkan extension, and applications that [“Acquire”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaymanager.tryacquiretarget) a [“DisplayTarget”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaytarget) using a [“WinRT”](https://docs.microsoft.com/en-us/uwp/api/) command such as [“winrt::Windows::Devices::Display::Core::DisplayManager.TryAcquireTarget()”](https://docs.microsoft.com/en-us/uwp/api/windows.devices.display.core.displaymanager.tryacquiretarget).

When control is acquired the application has exclusive access to the display until control is released or the application terminates. An application’s attempt to acquire is denied if a different application has already acquired the display.

## [](#_new_commands)New Commands

- [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireWinrtDisplayNV.html)
- [vkGetWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetWinrtDisplayNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_ACQUIRE_WINRT_DISPLAY_EXTENSION_NAME`
- `VK_NV_ACQUIRE_WINRT_DISPLAY_SPEC_VERSION`

## [](#_issues)Issues

1\) What should the platform substring be for this extension:

**RESOLVED**: The platform substring is “Winrt”.

The substring “Winrt” matches the fact that the OS API exposing the acquire and release functionality is called “WinRT”.

The substring “Win32” is wrong because the related “WinRT” API is explicitly **not** a “Win32” API. “WinRT” is a competing API family to the “Win32” API family.

The substring “Windows” is suboptimal because there could be more than one relevant API on the Windows platform. There is preference to use the more-specific substring “Winrt”.

2\) Should [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireWinrtDisplayNV.html) take a winRT DisplayTarget, or a Vulkan display handle as input?

**RESOLVED**: A Vulkan display handle. This matches the design of [vkAcquireXlibDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireXlibDisplayEXT.html).

3\) Should the acquire command be platform-independent named “vkAcquireDisplayNV”, or platform-specific named “vkAcquireWinrtDisplayNV”?

**RESOLVED**: Add a platform-specific command.

The inputs to the Acquire command are all Vulkan types. None are WinRT types. This opens the possibility of the winrt extension defining a platform-independent acquire command.

The X11 acquire command does need to accept a platform-specific parameter. This could be handled by adding to a platform-independent acquire command a params structure to which platform-dependent types can be chained by `pNext` pointer.

The prevailing opinion is that it would be odd to create a second platform-independent function that is used on the Windows 10 platform, but that is not used for the X11 platform. Since a Windows 10 platform-specific command is needed anyway for converting between vkDisplayKHR and platform-native handles, opinion was to create a platform-specific acquire function.

4\) Should the [vkGetWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetWinrtDisplayNV.html) parameter identifying a display be named “deviceRelativeId” or “adapterRelativeId”?

**RESOLVED**: The WinRT name is “AdapterRelativeId”. The name “adapter” is the Windows analog to a Vulkan “physical device”. Vulkan already has precedent to use the name `deviceLUID` for the concept that Windows APIs call “AdapterLuid”. Keeping form with this precedent, the name “deviceRelativeId” is chosen.

5\) Does [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireWinrtDisplayNV.html) cause the Windows desktop compositor to release a display?

**RESOLVED**: No. [vkAcquireWinrtDisplayNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireWinrtDisplayNV.html) does not itself cause the Windows desktop compositor to release a display. This action must be performed outside of Vulkan.

Beginning with Windows 10 version 2004 it is possible to cause the Windows desktop compositor to release a display by using the “Advanced display settings” sub-page of the “Display settings” control panel. See [https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors](https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors)

6\) Where can one find additional information about custom compositors for Windows 10?

**RESOLVED**: Relevant references are as follows.

According to Microsoft’s documentation on ["building a custom compositor"](https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors-compositor), the ability to write a custom compositor is not a replacement for a fullscreen desktop window. The feature is for writing compositor apps that drive specialized hardware.

Only certain editions of Windows 10 support custom compositors, ["documented here"](https://docs.microsoft.com/en-us/windows-hardware/drivers/display/specialized-monitors#windows-10-version-2004). The product type can be queried from Windows 10. See [https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo](https://docs.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo)

## [](#_version_history)Version History

- Revision 1, 2020-09-29 (Jeff Juliano)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_acquire_winrt_display).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0