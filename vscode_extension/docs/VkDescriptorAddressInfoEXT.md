# VkDescriptorAddressInfoEXT(3) Manual Page

## Name

VkDescriptorAddressInfoEXT - Structure specifying descriptor buffer address info



## [](#_c_specification)C Specification

Data describing a `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, or `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor is passed in a `VkDescriptorAddressInfoEXT` structure:

```c++
// Provided by VK_EXT_descriptor_buffer
typedef struct VkDescriptorAddressInfoEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceAddress    address;
    VkDeviceSize       range;
    VkFormat           format;
} VkDescriptorAddressInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `address` is either `0` or a device address at an offset in a buffer, where the base address can be queried from [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html).
- `range` is the size in bytes of the buffer or buffer view used by the descriptor.
- `format` is the format of the data elements in the buffer view and is ignored for buffers.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorAddressInfoEXT-None-09508)VUID-VkDescriptorAddressInfoEXT-None-09508  
  If `address` is not zero, and the descriptor is of type `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, then `format` **must** not be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkDescriptorAddressInfoEXT-address-08043)VUID-VkDescriptorAddressInfoEXT-address-08043  
  If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `address` **must** not be zero
- [](#VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08938)VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08938  
  If `address` is zero, `range` **must** be `VK_WHOLE_SIZE`
- [](#VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08939)VUID-VkDescriptorAddressInfoEXT-nullDescriptor-08939  
  If `address` is not zero, `range` **must** not be `VK_WHOLE_SIZE`
- [](#VUID-VkDescriptorAddressInfoEXT-range-08045)VUID-VkDescriptorAddressInfoEXT-range-08045  
  `range` **must** be less than or equal to the size of the buffer containing `address` minus the offset of `address` from the base address of the buffer
- [](#VUID-VkDescriptorAddressInfoEXT-range-08940)VUID-VkDescriptorAddressInfoEXT-range-08940  
  `range` **must** not be zero

Valid Usage (Implicit)

- [](#VUID-VkDescriptorAddressInfoEXT-sType-sType)VUID-VkDescriptorAddressInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_ADDRESS_INFO_EXT`
- [](#VUID-VkDescriptorAddressInfoEXT-pNext-pNext)VUID-VkDescriptorAddressInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDescriptorAddressInfoEXT-address-parameter)VUID-VkDescriptorAddressInfoEXT-address-parameter  
  If `address` is not `0`, `address` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-VkDescriptorAddressInfoEXT-format-parameter)VUID-VkDescriptorAddressInfoEXT-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, `address` **can** be zero. Loads from a null descriptor return zero values and stores and atomics to a null descriptor are discarded.

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorDataEXT.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorAddressInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0