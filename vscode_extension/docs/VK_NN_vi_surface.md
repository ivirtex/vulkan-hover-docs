# VK\_NN\_vi\_surface(3) Manual Page

## Name

VK\_NN\_vi\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

63

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Mathias Heyer \[GitLab]mheyer

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-12-02

**IP Status**

No known IP claims.

**Contributors**

- Mathias Heyer, NVIDIA
- Michael Chock, NVIDIA
- Yasuhiro Yoshioka, Nintendo
- Daniel Koch, NVIDIA

## [](#_description)Description

The `VK_NN_vi_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) associated with an `nn`::`vi`::`Layer`.

## [](#_new_commands)New Commands

- [vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateViSurfaceNN.html)

## [](#_new_structures)New Structures

- [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViSurfaceCreateInfoNN.html)

## [](#_new_bitmasks)New Bitmasks

- [VkViSurfaceCreateFlagsNN](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViSurfaceCreateFlagsNN.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NN_VI_SURFACE_EXTENSION_NAME`
- `VK_NN_VI_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_VI_SURFACE_CREATE_INFO_NN`

## [](#_issues)Issues

1\) Does VI need a way to query for compatibility between a particular physical device (and queue family?) and a specific VI display?

**RESOLVED**: No. It is currently always assumed that the device and display will always be compatible.

2\) [VkViSurfaceCreateInfoNN](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViSurfaceCreateInfoNN.html)::`window` is intended to store an `nn`::`vi`::`NativeWindowHandle`, but its declared type is a bare `void*` to store the window handle. Why the discrepancy?

**RESOLVED**: It is for C compatibility. The definition for the VI native window handle type is defined inside the `nn`::`vi` C++ namespace. This prevents its use in C source files. `nn`::`vi`::`NativeWindowHandle` is always defined to be `void*`, so this extension uses `void*` to match.

## [](#_version_history)Version History

- Revision 1, 2016-12-2 (Michael Chock)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NN_vi_surface).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0