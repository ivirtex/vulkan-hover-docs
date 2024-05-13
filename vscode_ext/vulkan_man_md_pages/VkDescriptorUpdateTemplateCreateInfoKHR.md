# VkDescriptorUpdateTemplateCreateInfo(3) Manual Page

## Name

VkDescriptorUpdateTemplateCreateInfo - Structure specifying parameters
of a newly created descriptor update template



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html)
structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_descriptor_update_template
typedef VkDescriptorUpdateTemplateCreateInfo VkDescriptorUpdateTemplateCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `descriptorUpdateEntryCount` is the number of elements in the
  `pDescriptorUpdateEntries` array.

- `pDescriptorUpdateEntries` is a pointer to an array of
  [VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntry.html)
  structures describing the descriptors to be updated by the descriptor
  update template.

- `templateType` Specifies the type of the descriptor update template.
  If set to `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET` it
  **can** only be used to update descriptor sets with a fixed
  `descriptorSetLayout`. If set to
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR` it **can**
  only be used to push descriptor sets using the provided
  `pipelineBindPoint`, `pipelineLayout`, and `set` number.

- `descriptorSetLayout` is the descriptor set layout used to build the
  descriptor update template. All descriptor sets which are going to be
  updated through the newly created descriptor update template **must**
  be created with a layout that matches (is the same as, or defined
  identically to) this layout. This parameter is ignored if
  `templateType` is not
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET`.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) indicating the type of
  the pipeline that will use the descriptors. This parameter is ignored
  if `templateType` is not
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

- `pipelineLayout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object
  used to program the bindings. This parameter is ignored if
  `templateType` is not
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

- `set` is the set number of the descriptor set in the pipeline layout
  that will be updated. This parameter is ignored if `templateType` is
  not `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00350"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00350"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00350  
  If `templateType` is
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET`,
  `descriptorSetLayout` **must** be a valid `VkDescriptorSetLayout`
  handle

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00351"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00351"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00351  
  If `templateType` is
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`,
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00352"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00352"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00352  
  If `templateType` is
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`,
  `pipelineLayout` **must** be a valid `VkPipelineLayout` handle

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00353"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00353"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-00353  
  If `templateType` is
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`, `set`
  **must** be the unique set number in the pipeline layout that uses a
  descriptor set layout that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-04615"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-04615"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-04615  
  If `templateType` is
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET`,
  `descriptorSetLayout` **must** not contain a binding with type
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-sType-sType"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-sType-sType"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO`

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-pNext-pNext"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-pNext-pNext"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-flags-zerobitmask"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-flags-zerobitmask"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkDescriptorUpdateTemplateCreateInfo-pDescriptorUpdateEntries-parameter"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-pDescriptorUpdateEntries-parameter"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-pDescriptorUpdateEntries-parameter  
  `pDescriptorUpdateEntries` **must** be a valid pointer to an array of
  `descriptorUpdateEntryCount` valid
  [VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntry.html)
  structures

- <a
  href="#VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-parameter"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-parameter"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-templateType-parameter  
  `templateType` **must** be a valid
  [VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html)
  value

- <a
  href="#VUID-VkDescriptorUpdateTemplateCreateInfo-descriptorUpdateEntryCount-arraylength"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-descriptorUpdateEntryCount-arraylength"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-descriptorUpdateEntryCount-arraylength  
  `descriptorUpdateEntryCount` **must** be greater than `0`

- <a href="#VUID-VkDescriptorUpdateTemplateCreateInfo-commonparent"
  id="VUID-VkDescriptorUpdateTemplateCreateInfo-commonparent"></a>
  VUID-VkDescriptorUpdateTemplateCreateInfo-commonparent  
  Both of `descriptorSetLayout`, and `pipelineLayout` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkDescriptorUpdateTemplateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateFlags.html),
[VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntry.html),
[VkDescriptorUpdateTemplateType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateType.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html),
[vkCreateDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplateKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorUpdateTemplateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
