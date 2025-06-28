# VK\_EXT\_acquire\_xlib\_display(3) Manual Page

## Name

VK\_EXT\_acquire\_xlib\_display - instance extension



## [](#_registered_extension_number)Registered Extension Number

90

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_direct\_mode\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_direct_mode_display.html)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_acquire_xlib_display%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_acquire_xlib_display%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

This extension allows an application to take exclusive control on a display currently associated with an X11 screen. When control is acquired, the display will be deassociated from the X11 screen until control is released or the specified display connection is closed. Essentially, the X11 screen will behave as if the monitor has been unplugged until control is released.

## [](#_new_commands)New Commands

- [vkAcquireXlibDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireXlibDisplayEXT.html)
- [vkGetRandROutputDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRandROutputDisplayEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_ACQUIRE_XLIB_DISPLAY_EXTENSION_NAME`
- `VK_EXT_ACQUIRE_XLIB_DISPLAY_SPEC_VERSION`

## [](#_issues)Issues

1\) Should [vkAcquireXlibDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireXlibDisplayEXT.html) take an RandR display ID, or a Vulkan display handle as input?

**RESOLVED**: A Vulkan display handle. Otherwise there would be no way to specify handles to displays that had been prevented from being included in the X11 display list by some native platform or vendor-specific mechanism.

2\) How does an application figure out which RandR display corresponds to a Vulkan display?

**RESOLVED**: A new function, [vkGetRandROutputDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRandROutputDisplayEXT.html), is introduced for this purpose.

3\) Should [vkGetRandROutputDisplayEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRandROutputDisplayEXT.html) be part of this extension, or a general Vulkan / RandR or Vulkan / Xlib extension?

**RESOLVED**: To avoid yet another extension, include it in this extension.

## [](#_version_history)Version History

- Revision 1, 2016-12-13 (James Jones)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_acquire_xlib_display)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0