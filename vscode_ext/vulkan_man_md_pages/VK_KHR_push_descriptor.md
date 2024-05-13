# VK_KHR_push_descriptor(3) Manual Page

## Name

VK_KHR_push_descriptor - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

81

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_1

- Interacts with VK_KHR_descriptor_update_template

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_push_descriptor%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_push_descriptor%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-12

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Michael Worcester, Imagination Technologies

## <a href="#_description" class="anchor"></a>Description

This extension allows descriptors to be written into the command buffer,
while the implementation is responsible for managing their memory. Push
descriptors may enable easier porting from older APIs and in some cases
can be more efficient than writing descriptors into descriptor sets.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetKHR.html)

If
[VK_KHR_descriptor_update_template](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_descriptor_update_template.html)
is supported:

- [vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)

If [Version 1.1](#versions-1.1) is supported:

- [vkCmdPushDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplateKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDevicePushDescriptorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePushDescriptorPropertiesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_PUSH_DESCRIPTOR_EXTENSION_NAME`

- `VK_KHR_PUSH_DESCRIPTOR_SPEC_VERSION`

- Extending
  [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html):

  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PUSH_DESCRIPTOR_PROPERTIES_KHR`

If
[VK_KHR_descriptor_update_template](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_descriptor_update_template.html)
is supported:

- Extending
  [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html):

  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

If [Version 1.1](#versions-1.1) is supported:

- Extending
  [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html):

  - `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-15 (Jeff Bolz)

  - Internal revisions

- Revision 2, 2017-09-12 (Tobias Hector)

  - Added interactions with Vulkan 1.1

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_push_descriptor"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
