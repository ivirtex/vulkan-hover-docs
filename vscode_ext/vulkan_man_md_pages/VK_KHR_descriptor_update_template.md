# VK_KHR_descriptor_update_template(3) Manual Page

## Name

VK_KHR_descriptor_update_template - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

86

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_EXT_debug_report

- Interacts with VK_KHR_push_descriptor

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Markus Tavenrath <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_descriptor_update_template%5D%20@mtavenrath%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_descriptor_update_template%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mtavenrath</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Interacts with [`VK_KHR_push_descriptor`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html)

**Contributors**  
- Jeff Bolz, NVIDIA

- Michael Worcester, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

Applications may wish to update a fixed set of descriptors in a large
number of descriptor sets very frequently, i.e. during initialization
phase or if it is required to rebuild descriptor sets for each frame.
For those cases it is also not unlikely that all information required to
update a single descriptor set is stored in a single struct. This
extension provides a way to update a fixed set of descriptors in a
single [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) with a pointer to a user
defined data structure describing the new descriptors.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

[vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)
is included as an interaction with
[`VK_KHR_push_descriptor`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html). If Vulkan 1.1
and `VK_KHR_push_descriptor` are supported, this is included by
[`VK_KHR_push_descriptor`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html).

The base functionality in this extension is included in core Vulkan 1.1,
with the KHR suffix omitted. The original type, enum and command names
are still available as aliases of the core functionality.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateKHR.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplateKHR.html)

- [vkDestroyDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplateKHR.html)

- [vkUpdateDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSetWithTemplateKHR.html)

If [VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html) is supported:

- [vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDescriptorUpdateTemplateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfoKHR.html)

- [VkDescriptorUpdateTemplateEntryKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntryKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkDescriptorUpdateTemplateTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateTypeKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkDescriptorUpdateTemplateCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME`

- `VK_KHR_DESCRIPTOR_UPDATE_TEMPLATE_SPEC_VERSION`

- Extending
  [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html):

  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO_KHR`

If [VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html) is supported:

- Extending
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html):

  - `VK_DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR_EXT`

If [VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html) is supported:

- Extending
  [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html):

  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-01-11 (Markus Tavenrath)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_descriptor_update_template"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
