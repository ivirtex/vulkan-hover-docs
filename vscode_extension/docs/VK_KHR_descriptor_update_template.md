# VK\_KHR\_descriptor\_update\_template(3) Manual Page

## Name

VK\_KHR\_descriptor\_update\_template - device extension



## [](#_registered_extension_number)Registered Extension Number

86

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_debug\_report
- Interacts with VK\_KHR\_push\_descriptor

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Markus Tavenrath [\[GitHub\]mtavenrath](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_descriptor_update_template%5D%20%40mtavenrath%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_descriptor_update_template%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Interacts with `VK_KHR_push_descriptor`

**Contributors**

- Jeff Bolz, NVIDIA
- Michael Worcester, Imagination Technologies

## [](#_description)Description

Applications may wish to update a fixed set of descriptors in a large number of descriptor sets very frequently, i.e. during initialization phase or if it is required to rebuild descriptor sets for each frame. For those cases it is also not unlikely that all information required to update a single descriptor set is stored in a single struct. This extension provides a way to update a fixed set of descriptors in a single [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) with a pointer to an application-defined data structure describing the new descriptors.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

[vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html) is included as an interaction with `VK_KHR_push_descriptor`. If Vulkan 1.1 and `VK_KHR_push_descriptor` are supported, this is included by `VK_KHR_push_descriptor`.

The base functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_object_types)New Object Types

- [VkDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateKHR.html)

## [](#_new_commands)New Commands

- [vkCreateDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDescriptorUpdateTemplateKHR.html)
- [vkDestroyDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyDescriptorUpdateTemplateKHR.html)
- [vkUpdateDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplateKHR.html)

If [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) is supported:

- [vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)

## [](#_new_structures)New Structures

- [VkDescriptorUpdateTemplateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfoKHR.html)
- [VkDescriptorUpdateTemplateEntryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateEntryKHR.html)

## [](#_new_enums)New Enums

- [VkDescriptorUpdateTemplateTypeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateTypeKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkDescriptorUpdateTemplateCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME`
- `VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_SPEC_VERSION`
- Extending [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateType.html):
  
  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR`

If [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html):
  
  - `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR_EXT`

If [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) is supported:

- Extending [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateType.html):
  
  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

## [](#_version_history)Version History

- Revision 1, 2016-01-11 (Markus Tavenrath)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_descriptor_update_template).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0