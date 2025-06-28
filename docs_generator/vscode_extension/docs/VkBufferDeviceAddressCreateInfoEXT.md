# VkBufferDeviceAddressCreateInfoEXT(3) Manual Page

## Name

VkBufferDeviceAddressCreateInfoEXT - Request a specific address for a buffer



## [](#_c_specification)C Specification

Alternatively, to request a specific device address for a buffer, add a [VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressCreateInfoEXT.html) structure to the `pNext` chain of the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structure. The `VkBufferDeviceAddressCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_buffer_device_address
typedef struct VkBufferDeviceAddressCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceAddress    deviceAddress;
} VkBufferDeviceAddressCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `deviceAddress` is the device address requested for the buffer.

## [](#_description)Description

If `deviceAddress` is zero, no specific address is requested.

If `deviceAddress` is not zero, then it **must** be an address retrieved from an identically created buffer on the same implementation. The buffer **must** also be bound to an identically created `VkDeviceMemory` object.

If this structure is not present, it is as if `deviceAddress` is zero.

Applications **should** avoid creating buffers with application-provided addresses and implementation-provided addresses in the same process, to reduce the likelihood of `VK_ERROR_INVALID_DEVICE_ADDRESS_EXT` errors.

Valid Usage (Implicit)

- [](#VUID-VkBufferDeviceAddressCreateInfoEXT-sType-sType)VUID-VkBufferDeviceAddressCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_buffer_device_address.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferDeviceAddressCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0