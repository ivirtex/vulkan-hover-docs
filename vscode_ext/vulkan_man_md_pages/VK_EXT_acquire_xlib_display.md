# VK_EXT_acquire_xlib_display(3) Manual Page

## Name

VK_EXT_acquire_xlib_display - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

90

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_direct_mode_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_direct_mode_display.html)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_acquire_xlib_display%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_acquire_xlib_display%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-12-13

**IP Status**  
No known IP claims.

**Contributors**  
- Dave Airlie, Red Hat

- Pierre Boudier, NVIDIA

- James Jones, NVIDIA

- Damien Leone, NVIDIA

- Pierre-Loup Griffais, Valve

- Liam Middlebrook, NVIDIA

- Daniel Vetter, Intel

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to take exclusive control on a
display currently associated with an X11 screen. When control is
acquired, the display will be deassociated from the X11 screen until
control is released or the specified display connection is closed.
Essentially, the X11 screen will behave as if the monitor has been
unplugged until control is released.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkAcquireXlibDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireXlibDisplayEXT.html)

- [vkGetRandROutputDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRandROutputDisplayEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_ACQUIRE_XLIB_DISPLAY_EXTENSION_NAME`

- `VK_EXT_ACQUIRE_XLIB_DISPLAY_SPEC_VERSION`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should [vkAcquireXlibDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireXlibDisplayEXT.html) take
an RandR display ID, or a Vulkan display handle as input?

**RESOLVED**: A Vulkan display handle. Otherwise there would be no way
to specify handles to displays that had been prevented from being
included in the X11 display list by some native platform or
vendor-specific mechanism.

2\) How does an application figure out which RandR display corresponds
to a Vulkan display?

**RESOLVED**: A new function,
[vkGetRandROutputDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRandROutputDisplayEXT.html), is
introduced for this purpose.

3\) Should [vkGetRandROutputDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRandROutputDisplayEXT.html)
be part of this extension, or a general Vulkan / RandR or Vulkan / Xlib
extension?

**RESOLVED**: To avoid yet another extension, include it in this
extension.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-13 (James Jones)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_acquire_xlib_display"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
