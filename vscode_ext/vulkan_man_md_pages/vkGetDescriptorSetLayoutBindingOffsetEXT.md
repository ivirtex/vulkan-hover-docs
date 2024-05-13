# vkGetDescriptorSetLayoutBindingOffsetEXT(3) Manual Page

## Name

vkGetDescriptorSetLayoutBindingOffsetEXT - Get the offset of a binding
within a descriptor set layout



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the offset of a binding within a descriptor set layout in memory,
call:

``` c
// Provided by VK_EXT_descriptor_buffer
void vkGetDescriptorSetLayoutBindingOffsetEXT(
    VkDevice                                    device,
    VkDescriptorSetLayout                       layout,
    uint32_t                                    binding,
    VkDeviceSize*                               pOffset);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that gets the offset.

- `layout` is the descriptor set layout being queried.

- `binding` is the binding number being queried.

- `pOffset` is a pointer to [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) where the
  byte offset of the binding will be written.

## <a href="#_description" class="anchor"></a>Description

Each binding in a descriptor set layout is assigned an offset in memory
by the implementation. When a shader accesses a resource with that
binding, it will access the bound descriptor buffer from that offset to
look for its descriptor. This command provides an application with that
offset, so that descriptors can be placed in the correct locations. The
precise location accessed by a shader for a given descriptor is as
follows:

  
location = bufferAddress + setOffset + descriptorOffset + (arrayElement
× descriptorSize)

where bufferAddress and setOffset are the base address and offset for
the identified descriptor set as specified by
[vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBuffersEXT.html) and
[vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html),
descriptorOffset is the offset for the binding returned by this command,
arrayElement is the index into the array specified in the shader, and
descriptorSize is the size of the relevant descriptor as obtained from
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html).
Applications are responsible for placing valid descriptors at the
expected location in order for a shader to access it. The overall offset
added to bufferAddress to calculate location **must** be less than
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxSamplerDescriptorBufferRange`
for samplers and
[VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxResourceDescriptorBufferRange`
for resources.

If any `binding` in `layout` is
`VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, that `binding`
**must** have the largest offset of any `binding`.

A descriptor `binding` with type `VK_DESCRIPTOR_TYPE_MUTABLE_VALVE`
**can** be used. Any potential types in
[VkMutableDescriptorTypeCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoVALVE.html)::`pDescriptorTypes`
for `binding` share the same offset. If the size of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-mutable"
target="_blank" rel="noopener">mutable descriptor</a> is larger than the
size of a concrete descriptor type being accessed, the padding area is
ignored by the implementation.

Valid Usage

- <a href="#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-None-08013"
  id="VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-None-08013"></a>
  VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-None-08013  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-08014"
  id="VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-08014"></a>
  VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-08014  
  `layout` **must** have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` flag set

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-device-parameter"
  id="VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-device-parameter"></a>
  VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parameter"
  id="VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parameter"></a>
  VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parameter  
  `layout` **must** be a valid
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handle

- <a
  href="#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-pOffset-parameter"
  id="VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-pOffset-parameter"></a>
  VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-pOffset-parameter  
  `pOffset` **must** be a valid pointer to a
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) value

- <a href="#VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parent"
  id="VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parent"></a>
  VUID-vkGetDescriptorSetLayoutBindingOffsetEXT-layout-parent  
  `layout` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDescriptorSetLayoutBindingOffsetEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
