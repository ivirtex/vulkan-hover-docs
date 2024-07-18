# vkGetDescriptorSetLayoutSizeEXT(3) Manual Page

## Name

vkGetDescriptorSetLayoutSizeEXT - Get the size of a descriptor set
layout in memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the amount of memory needed to store all descriptors with a
given layout, call:

``` c
// Provided by VK_EXT_descriptor_buffer
void vkGetDescriptorSetLayoutSizeEXT(
    VkDevice                                    device,
    VkDescriptorSetLayout                       layout,
    VkDeviceSize*                               pLayoutSizeInBytes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that gets the size.

- `layout` is the descriptor set layout being queried.

- `pLayoutSizeInBytes` is a pointer to [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)
  where the size in bytes will be written.

## <a href="#_description" class="anchor"></a>Description

The size of a descriptor set layout will be at least as large as the sum
total of the size of all descriptors in the layout, and **may** be
larger. This size represents the amount of memory that will be required
to store all of the descriptors for this layout in memory, when placed
according to the layout’s offsets as obtained by
[vkGetDescriptorSetLayoutBindingOffsetEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutBindingOffsetEXT.html).

If any `binding` in `layout` is
`VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, the returned size
includes space for the maximum `descriptorCount` descriptors as declared
for that `binding`. To compute the required size of a descriptor set
with a `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`:

  
size = offset + descriptorSize × variableDescriptorCount

where offset is obtained by
[vkGetDescriptorSetLayoutBindingOffsetEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutBindingOffsetEXT.html)
and descriptorSize is the size of the relevant descriptor as obtained
from
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html),
and variableDescriptorCount is the equivalent of
[VkDescriptorSetVariableDescriptorCountAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountAllocateInfo.html)::`pDescriptorCounts`.
For `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, variableDescriptorCount
is the size in bytes for the inline uniform block, and descriptorSize is
1.

If
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`combinedImageSamplerDescriptorSingleArray`
is `VK_FALSE` and the variable descriptor type is
`VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, variableDescriptorCount is
always considered to be the upper bound.

Valid Usage

- <a href="#VUID-vkGetDescriptorSetLayoutSizeEXT-None-08011"
  id="VUID-vkGetDescriptorSetLayoutSizeEXT-None-08011"></a>
  VUID-vkGetDescriptorSetLayoutSizeEXT-None-08011  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetDescriptorSetLayoutSizeEXT-layout-08012"
  id="VUID-vkGetDescriptorSetLayoutSizeEXT-layout-08012"></a>
  VUID-vkGetDescriptorSetLayoutSizeEXT-layout-08012  
  `layout` **must** have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` flag set

Valid Usage (Implicit)

- <a href="#VUID-vkGetDescriptorSetLayoutSizeEXT-device-parameter"
  id="VUID-vkGetDescriptorSetLayoutSizeEXT-device-parameter"></a>
  VUID-vkGetDescriptorSetLayoutSizeEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parameter"
  id="VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parameter"></a>
  VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parameter  
  `layout` **must** be a valid
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handle

- <a
  href="#VUID-vkGetDescriptorSetLayoutSizeEXT-pLayoutSizeInBytes-parameter"
  id="VUID-vkGetDescriptorSetLayoutSizeEXT-pLayoutSizeInBytes-parameter"></a>
  VUID-vkGetDescriptorSetLayoutSizeEXT-pLayoutSizeInBytes-parameter  
  `pLayoutSizeInBytes` **must** be a valid pointer to a
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) value

- <a href="#VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parent"
  id="VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parent"></a>
  VUID-vkGetDescriptorSetLayoutSizeEXT-layout-parent  
  `layout` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDescriptorSetLayoutSizeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
