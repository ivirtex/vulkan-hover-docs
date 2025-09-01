# VkDeviceMemoryOpaqueCaptureAddressInfo(3) Manual Page

## Name

VkDeviceMemoryOpaqueCaptureAddressInfo - Structure specifying the memory object to query an address for



## [](#_c_specification)C Specification

The `VkDeviceMemoryOpaqueCaptureAddressInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkDeviceMemoryOpaqueCaptureAddressInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
} VkDeviceMemoryOpaqueCaptureAddressInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_buffer_device_address
typedef VkDeviceMemoryOpaqueCaptureAddressInfo VkDeviceMemoryOpaqueCaptureAddressInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` specifies the memory whose address is being queried.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-03336)VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-03336  
  `memory` **must** have been allocated with `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT`

Valid Usage (Implicit)

- [](#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-sType-sType)VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO`
- [](#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-pNext-pNext)VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-parameter)VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

## [](#_see_also)See Also

[VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html), [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceMemoryOpaqueCaptureAddressInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0