# VK\_EXT\_vertex\_attribute\_robustness(3) Manual Page

## Name

VK\_EXT\_vertex\_attribute\_robustness - device extension



## [](#_registered_extension_number)Registered Extension Number

609

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_maintenance9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance9.html) extension

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_vertex_attribute_robustness%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_vertex_attribute_robustness%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-11-01

**IP Status**

No known IP claims.

**Contributors**

- Daniel Story, Nintendo

## [](#_description)Description

It can be detrimental to performance for applications to have to define fake vertex attribute locations and buffer bindings for vertex shaders that may reference attribute locations for which there is no vertex data.

This extension allows applications to not have to specify fake vertex attribute locations, and if the vertex shader reads those attributes it will read (0,0,0,0) or (0,0,0,1).

## [](#_promotion_to_vk_khr_maintenance9)Promotion to `VK_KHR_maintenance9`

The same functionality is provided by [VK\_KHR\_maintenance9](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance9.html), but enabled by the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature instead. The [VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT.html) structure was not included in the maintenance extension.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexAttributeRobustnessFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_VERTEX_ATTRIBUTE_ROBUSTNESS_EXTENSION_NAME`
- `VK_EXT_VERTEX_ATTRIBUTE_ROBUSTNESS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_ROBUSTNESS_FEATURES_EXT`

## [](#_issues)Issues

None

## [](#_version_history)Version History

- Revision 1, 2024-11-01 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_vertex_attribute_robustness).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0