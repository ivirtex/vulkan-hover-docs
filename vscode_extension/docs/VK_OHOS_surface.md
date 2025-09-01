# VK\_OHOS\_surface(3) Manual Page

## Name

VK\_OHOS\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

588

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Weilan Chen [\[GitHub\]wchen-h](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_OHOS_surface%5D%20%40wchen-h%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_OHOS_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-05-16

**IP Status**

No known IP claims.

**Contributors**

- Weilan Chen, Huawei
- Zehui Lin, Huawei
- Pan Gao, Huawei
- Zhao Zhang, Huawei
- Yang Shi, Huawei

## [](#_description)Description

The `VK_OHOS_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to an [OHNativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/OHNativeWindow.html), the native surface type of Open Harmony OS. Common consumer endpoints for `OHNativeWindows` are the system window compositor, video encoders, and application-specific compositors.

## [](#_new_base_types)New Base Types

- [OHNativeWindow](https://registry.khronos.org/vulkan/specs/latest/man/html/OHNativeWindow.html)

## [](#_new_commands)New Commands

- [vkCreateSurfaceOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSurfaceOHOS.html)

## [](#_new_structures)New Structures

- [VkSurfaceCreateInfoOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCreateInfoOHOS.html)

## [](#_new_bitmasks)New Bitmasks

- [VkSurfaceCreateFlagsOHOS](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCreateFlagsOHOS.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_OHOS_SURFACE_EXTENSION_NAME`
- `VK_OHOS_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_SURFACE_CREATE_INFO_OHOS`

## [](#_version_history)Version History

- Revision 1, 2025-05-19 (Weilan Chen)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_OHOS_surface).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0