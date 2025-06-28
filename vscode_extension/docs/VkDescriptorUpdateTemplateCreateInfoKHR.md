# VkDescriptorUpdateTemplateCreateInfo(3) Manual Page

## Name

VkDescriptorUpdateTemplateCreateInfo - Structure specifying parameters of a newly created descriptor update template



## [](#_c_specification)C Specification

The [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfo.html) structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDescriptorUpdateTemplateCreateInfo {
    VkStructureType                           sType;
    const void*                               pNext;
    VkDescriptorUpdateTemplateCreateFlags     flags;
    uint32_t                                  descriptorUpdateEntryCount;
    const VkDescriptorUpdateTemplateEntry*    pDescriptorUpdateEntries;
    VkDescriptorUpdateTemplateType            templateType;
    VkDescriptorSetLayout                     descriptorSetLayout;
    VkPipelineBindPoint                       pipelineBindPoint;
    VkPipelineLayout                          pipelineLayout;
    uint32_t                                  set;
} VkDescriptorUpdateTemplateCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_descriptor_update_template
typedef VkDescriptorUpdateTemplateCreateInfo VkDescriptorUpdateTemplateCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `descriptorUpdateEntryCount` is the number of elements in the `pDescriptorUpdateEntries` array.
- `pDescriptorUpdateEntries` is a pointer to an array of [VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateEntry.html) structures describing the descriptors to be updated by the descriptor update template.
- `templateType` Specifies the type of the descriptor update template. If set to `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET` it **can** only be used to update descriptor sets with a fixed `descriptorSetLayout`. If set to `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS` it **can** only be used to push descriptor sets using the provided `pipelineBindPoint`, `pipelineLayout`, and `set` number.
- `descriptorSetLayout` is the descriptor set layout used to build the descriptor update template. All descriptor sets which are going to be updated through the newly created descriptor update template **must** be created with a layout that matches (is the same as, or defined identically to) this layout. This parameter is ignored if `templateType` is not `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET`. The implementation **must** not access this object outside of the duration of the command this structure is passed to.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) indicating the type of the pipeline that will use the descriptors. This parameter is ignored if `templateType` is not `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`
- `pipelineLayout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html) object used to program the bindings. This parameter is ignored if `templateType` is not `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`
- `set` is the set number of the descriptor set in the pipeline layout that will be updated. This parameter is ignored if `templateType` is not `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00350)VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00350  
  If `templateType` is `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET`, `descriptorSetLayout` **must** be a valid `VkDescriptorSetLayout` handle
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-10355)VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-10355  
  If `templateType` is `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`, and the [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) extension is not enabled, [`pushDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pushDescriptor) **must** be enabled
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00351)VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00351  
  If `templateType` is `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`, `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00352)VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00352  
  If `templateType` is `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`, `pipelineLayout` **must** be a valid `VkPipelineLayout` handle
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00353)VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00353  
  If `templateType` is `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS`, `set` **must** be the unique set number in the pipeline layout that uses a descriptor set layout that was created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-04615)VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-04615  
  If `templateType` is `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET`, `descriptorSetLayout` **must** not contain a binding with type `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-sType-sType)VUID-VkDescriptorUpdateTemplateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO`
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-pNext-pNext)VUID-VkDescriptorUpdateTemplateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-flags-zerobitmask)VUID-VkDescriptorUpdateTemplateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-pDescriptorUpdateEntries-parameter)VUID-VkDescriptorUpdateTemplateCreateInfo-pDescriptorUpdateEntries-parameter  
  `pDescriptorUpdateEntries` **must** be a valid pointer to an array of `descriptorUpdateEntryCount` valid [VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateEntry.html) structures
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-parameter)VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-parameter  
  `templateType` **must** be a valid [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateType.html) value
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-descriptorUpdateEntryCount-arraylength)VUID-VkDescriptorUpdateTemplateCreateInfo-descriptorUpdateEntryCount-arraylength  
  `descriptorUpdateEntryCount` **must** be greater than `0`
- [](#VUID-VkDescriptorUpdateTemplateCreateInfo-commonparent)VUID-VkDescriptorUpdateTemplateCreateInfo-commonparent  
  Both of `descriptorSetLayout`, and `pipelineLayout` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkDescriptorUpdateTemplateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateFlags.html), [VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateEntry.html), [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateType.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDescriptorUpdateTemplate.html), [vkCreateDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDescriptorUpdateTemplateKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorUpdateTemplateCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0