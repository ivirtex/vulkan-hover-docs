# VkDescriptorAddressInfoEXT(3) Manual Page

## Name

VkDescriptorAddressInfoEXT - Structure specifying descriptor buffer
address info



## <a href="#_c_specification" class="anchor"></a>C Specification

Data describing a `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`,
`VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`,
`VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, or
`VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is passed in a
`VkDescriptorAddressInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorAddressInfoEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceAddress    address;
    VkDeviceSize       range;
    VkFormat           format;
} VkDescriptorAddressInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `address` is either `0` or a device address at an offset in a buffer,
  where the base address can be queried from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html).

- `range` is the size in bytes of the buffer or buffer view used by the
  descriptor.

- `format` is the format of the data elements in the buffer view and is
  ignored for buffers.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDescriptorAddressInfoEXT-None-09508"
  id="VUID-VkDescriptorAddressInfoEXT-None-09508"></a>
  VUID-VkDescriptorAddressInfoEXT-None-09508  
  If `address` is not zero, and the descriptor is of type
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, then `format` **must** not
  be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkDescriptorAddressInfoEXT-address-08043"
  id="VUID-VkDescriptorAddressInfoEXT-address-08043"></a>
  VUID-VkDescriptorAddressInfoEXT-address-08043  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, `address` **must** not be zero

- <a href="#VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08938"
  id="VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08938"></a>
  VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08938  
  If `address` is zero, `range` **must** be `VK_WHOLE_SIZE`

- <a href="#VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08939"
  id="VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08939"></a>
  VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08939  
  If `address` is not zero, `range` **must** not be `VK_WHOLE_SIZE`

- <a href="#VUID-VkDescriptorAddressInfoEXT-None-08044"
  id="VUID-VkDescriptorAddressInfoEXT-None-08044"></a>
  VUID-VkDescriptorAddressInfoEXT-None-08044  
  If `address` is not zero, `address` **must** be a valid device address
  at an offset within a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)

- <a href="#VUID-VkDescriptorAddressInfoEXT-range-08045"
  id="VUID-VkDescriptorAddressInfoEXT-range-08045"></a>
  VUID-VkDescriptorAddressInfoEXT-range-08045  
  `range` **must** be less than or equal to the size of the buffer
  containing `address` minus the offset of `address` from the base
  address of the buffer

- <a href="#VUID-VkDescriptorAddressInfoEXT-range-08940"
  id="VUID-VkDescriptorAddressInfoEXT-range-08940"></a>
  VUID-VkDescriptorAddressInfoEXT-range-08940  
  `range` **must** not be zero

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorAddressInfoEXT-sType-sType"
  id="VUID-VkDescriptorAddressInfoEXT-sType-sType"></a>
  VUID-VkDescriptorAddressInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_ADDRESS_INFO_EXT`

- <a href="#VUID-VkDescriptorAddressInfoEXT-pNext-pNext"
  id="VUID-VkDescriptorAddressInfoEXT-pNext-pNext"></a>
  VUID-VkDescriptorAddressInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDescriptorAddressInfoEXT-format-parameter"
  id="VUID-VkDescriptorAddressInfoEXT-format-parameter"></a>
  VUID-VkDescriptorAddressInfoEXT-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
is enabled, `address` **can** be zero. Loads from a null descriptor
return zero values and stores and atomics to a null descriptor are
discarded.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorDataEXT.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorAddressInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
