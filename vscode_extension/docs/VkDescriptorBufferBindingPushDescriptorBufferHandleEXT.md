# VkDescriptorBufferBindingPushDescriptorBufferHandleEXT(3) Manual Page

## Name

VkDescriptorBufferBindingPushDescriptorBufferHandleEXT - Structure
specifying push descriptor buffer binding information



## <a href="#_c_specification" class="anchor"></a>C Specification

When the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-bufferlessPushDescriptors"
target="_blank"
rel="noopener"><code>VkPhysicalDeviceDescriptorBufferPropertiesEXT</code>::<code>bufferlessPushDescriptors</code></a>
property is `VK_FALSE`, the `VkBuffer` handle of the buffer for push
descriptors is passed in a
`VkDescriptorBufferBindingPushDescriptorBufferHandleEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorBufferBindingPushDescriptorBufferHandleEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBuffer           buffer;
} VkDescriptorBufferBindingPushDescriptorBufferHandleEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` is the `VkBuffer` handle of the buffer for push descriptors.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-bufferlessPushDescriptors-08059"
  id="VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-bufferlessPushDescriptors-08059"></a>
  VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-bufferlessPushDescriptors-08059  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-bufferlessPushDescriptors"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDescriptorBufferPropertiesEXT</code>::<code>bufferlessPushDescriptors</code></a>
  **must** be `VK_FALSE`

Valid Usage (Implicit)

- <a
  href="#VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-sType-sType"
  id="VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-sType-sType"></a>
  VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_BUFFER_BINDING_PUSH_DESCRIPTOR_BUFFER_HANDLE_EXT`

- <a
  href="#VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-buffer-parameter"
  id="VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-buffer-parameter"></a>
  VUID-VkDescriptorBufferBindingPushDescriptorBufferHandleEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorBufferBindingPushDescriptorBufferHandleEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
