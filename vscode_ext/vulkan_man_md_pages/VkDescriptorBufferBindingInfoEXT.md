# VkDescriptorBufferBindingInfoEXT(3) Manual Page

## Name

VkDescriptorBufferBindingInfoEXT - Structure specifying descriptor
buffer binding information



## <a href="#_c_specification" class="anchor"></a>C Specification

Data describing a descriptor buffer binding is passed in a
`VkDescriptorBufferBindingInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorBufferBindingInfoEXT {
    VkStructureType       sType;
    void*                 pNext;
    VkDeviceAddress       address;
    VkBufferUsageFlags    usage;
} VkDescriptorBufferBindingInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `address` is a [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) specifying the
  device address defining the descriptor buffer to be bound.

- `usage` is a bitmask of
  [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html) specifying the
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`usage` for the buffer
  from which `address` was queried.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain includes a
[VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
structure,
[VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)::`usage`
from that structure is used instead of `usage` from this structure.

Valid Usage

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-None-09499"
  id="VUID-VkDescriptorBufferBindingInfoEXT-None-09499"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-None-09499  
  If the `pNext` chain does not include a
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
  structure, `usage` must be a valid combination of
  [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html) values

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-None-09500"
  id="VUID-VkDescriptorBufferBindingInfoEXT-None-09500"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-None-09500  
  If the `pNext` chain does not include a
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
  structure, `usage` must not be 0

- <a
  href="#VUID-VkDescriptorBufferBindingInfoEXT-bufferlessPushDescriptors-08056"
  id="VUID-VkDescriptorBufferBindingInfoEXT-bufferlessPushDescriptors-08056"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-bufferlessPushDescriptors-08056  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-bufferlessPushDescriptors"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDescriptorBufferPropertiesEXT</code>::<code>bufferlessPushDescriptors</code></a>
  is `VK_FALSE`, and `usage` contains
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`, then the
  `pNext` chain **must** include a
  [VkDescriptorBufferBindingPushDescriptorBufferHandleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferBindingPushDescriptorBufferHandleEXT.html)
  structure

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-address-08057"
  id="VUID-VkDescriptorBufferBindingInfoEXT-address-08057"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-address-08057  
  `address` **must** be aligned to
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferOffsetAlignment`

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-usage-08122"
  id="VUID-VkDescriptorBufferBindingInfoEXT-usage-08122"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-usage-08122  
  If `usage` includes
  `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`, `address`
  **must** be an address within a valid buffer that was created with
  `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-usage-08123"
  id="VUID-VkDescriptorBufferBindingInfoEXT-usage-08123"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-usage-08123  
  If `usage` includes
  `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`, `address`
  **must** be an address within a valid buffer that was created with
  `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-usage-08124"
  id="VUID-VkDescriptorBufferBindingInfoEXT-usage-08124"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-usage-08124  
  If `usage` includes
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`,
  `address` **must** be an address within a valid buffer that was
  created with
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-sType-sType"
  id="VUID-VkDescriptorBufferBindingInfoEXT-sType-sType"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_BUFFER_BINDING_INFO_EXT`

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-pNext-pNext"
  id="VUID-VkDescriptorBufferBindingInfoEXT-pNext-pNext"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkBufferUsageFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2CreateInfoKHR.html)
  or
  [VkDescriptorBufferBindingPushDescriptorBufferHandleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferBindingPushDescriptorBufferHandleEXT.html)

- <a href="#VUID-VkDescriptorBufferBindingInfoEXT-sType-unique"
  id="VUID-VkDescriptorBufferBindingInfoEXT-sType-unique"></a>
  VUID-VkDescriptorBufferBindingInfoEXT-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBuffersEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorBufferBindingInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
