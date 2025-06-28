# VK\_EXT\_multi\_draw(3) Manual Page

## Name

VK\_EXT\_multi\_draw - device extension



## [](#_registered_extension_number)Registered Extension Number

393

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_multi_draw%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_multi_draw%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-05-19

**Interactions and External Dependencies**

- Interacts with Vulkan 1.1.
- Interacts with `VK_KHR_shader_draw_parameters`.

**IP Status**

No known IP claims.

**Contributors**

- Mike Blumenkrantz, VALVE
- Piers Daniell, NVIDIA
- Faith Ekstrand, INTEL
- Spencer Fricke, SAMSUNG
- Ricardo Garcia, IGALIA
- Jon Leech, KHRONOS
- Stu Smith, AMD

## [](#_description)Description

Processing multiple draw commands in sequence incurs measurable overhead within drivers due to repeated state checks and updates during dispatch. This extension enables passing the entire sequence of draws directly to the driver in order to avoid any such overhead, using an array of [VkMultiDrawInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiDrawInfoEXT.html) or [VkMultiDrawIndexedInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiDrawIndexedInfoEXT.html) structs with `vkCmdDrawMultiEXT` or `vkCmdDrawMultiIndexedEXT`, respectively. These functions could be used any time multiple draw commands are being recorded without any state changes between them in order to maximize performance.

## [](#_new_commands)New Commands

- [vkCmdDrawMultiEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMultiEXT.html)
- [vkCmdDrawMultiIndexedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDrawMultiIndexedEXT.html)

## [](#_new_structures)New Structures

- [VkMultiDrawIndexedInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiDrawIndexedInfoEXT.html)
- [VkMultiDrawInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiDrawInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceMultiDrawFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultiDrawFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceMultiDrawPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMultiDrawPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_MULTI_DRAW_EXTENSION_NAME`
- `VK_EXT_MULTI_DRAW_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MULTI_DRAW_PROPERTIES_EXT`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- (modified)`DrawIndex`

## [](#_version_history)Version History

- Revision 1, 2021-01-20 (Mike Blumenkrantz)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_multi_draw)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0