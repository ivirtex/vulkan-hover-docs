# VkPushDescriptorSetWithTemplateInfoKHR(3) Manual Page

## Name

VkPushDescriptorSetWithTemplateInfoKHR - Structure specifying a
descriptor set push operation using a descriptor update template



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPushDescriptorSetWithTemplateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance6 with VK_KHR_push_descriptor
typedef struct VkPushDescriptorSetWithTemplateInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkDescriptorUpdateTemplate    descriptorUpdateTemplate;
    VkPipelineLayout              layout;
    uint32_t                      set;
    const void*                   pData;
} VkPushDescriptorSetWithTemplateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `descriptorUpdateTemplate` is a descriptor update template defining
  how to interpret the descriptor information in `pData`.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings. It **must** be compatible with the layout used
  to create the `descriptorUpdateTemplate` handle. If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicPipelineLayout"
  target="_blank" rel="noopener"><code>dynamicPipelineLayout</code></a>
  feature is enabled, `layout` **can** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the layout **must** be
  specified by chaining
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure off the `pNext`

- `set` is the set number of the descriptor set in the pipeline layout
  that will be updated. This **must** be the same number used to create
  the `descriptorUpdateTemplate` handle.

- `pData` is a pointer to memory containing descriptors for the
  templated update.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-commandBuffer-00366"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-commandBuffer-00366"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-commandBuffer-00366  
  The `pipelineBindPoint` specified during the creation of the
  descriptor update template **must** be supported by the
  `commandBuffer`’s parent `VkCommandPool`’s queue family

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-pData-01686"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-pData-01686"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-pData-01686  
  `pData` **must** be a valid pointer to a memory containing one or more
  valid instances of
  [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html),
  [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html), or
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) in a layout defined by
  `descriptorUpdateTemplate` when it was created with
  [vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html)

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-07993"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-07993"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-07993  
  `layout` **must** be compatible with the layout used to create
  `descriptorUpdateTemplate`

- <a
  href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-descriptorUpdateTemplate-07994"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-descriptorUpdateTemplate-07994"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-descriptorUpdateTemplate-07994  
  `descriptorUpdateTemplate` **must** have been created with a
  `templateType` of
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07995"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07995"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07995  
  `set` **must** be the same value used to create
  `descriptorUpdateTemplate`

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07304"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07304"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07304  
  `set` **must** be less than
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07305"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07305"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-set-07305  
  `set` **must** be the unique set number in the pipeline layout that
  uses a descriptor set layout that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`

<!-- -->

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-None-09495"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-None-09495"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout)
  feature is not enabled, `layout` **must** be a valid
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-09496"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-09496"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-09496  
  If `layout` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `pNext`
  chain **must** include a valid
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-sType-sType"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-sType-sType"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PUSH_DESCRIPTOR_SET_WITH_TEMPLATE_INFO_KHR`

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-pNext-pNext"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-pNext-pNext"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-sType-unique"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-sType-unique"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a
  href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-descriptorUpdateTemplate-parameter"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-descriptorUpdateTemplate-parameter"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-descriptorUpdateTemplate-parameter  
  `descriptorUpdateTemplate` **must** be a valid
  [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) handle

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-parameter"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-parameter"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-layout-parameter  
  If `layout` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `layout`
  **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-pData-parameter"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-pData-parameter"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-pData-parameter  
  `pData` **must** be a pointer value

- <a href="#VUID-VkPushDescriptorSetWithTemplateInfoKHR-commonparent"
  id="VUID-VkPushDescriptorSetWithTemplateInfoKHR-commonparent"></a>
  VUID-VkPushDescriptorSetWithTemplateInfoKHR-commonparent  
  Both of `descriptorUpdateTemplate`, and `layout` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html),
[VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdPushDescriptorSetWithTemplate2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetWithTemplate2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPushDescriptorSetWithTemplateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
