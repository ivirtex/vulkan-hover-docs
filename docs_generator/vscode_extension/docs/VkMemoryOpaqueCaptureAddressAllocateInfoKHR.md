# VkMemoryOpaqueCaptureAddressAllocateInfo(3) Manual Page

## Name

VkMemoryOpaqueCaptureAddressAllocateInfo - Request a specific address for a memory allocation



## [](#_c_specification)C Specification

To request a specific device address for a memory allocation, add a [VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html) structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure. The `VkMemoryOpaqueCaptureAddressAllocateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkMemoryOpaqueCaptureAddressAllocateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           opaqueCaptureAddress;
} VkMemoryOpaqueCaptureAddressAllocateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_buffer_device_address
typedef VkMemoryOpaqueCaptureAddressAllocateInfo VkMemoryOpaqueCaptureAddressAllocateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `opaqueCaptureAddress` is the opaque capture address requested for the memory allocation.

## [](#_description)Description

If `opaqueCaptureAddress` is zero, no specific address is requested.

If `opaqueCaptureAddress` is not zero, it **should** be an address retrieved from [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html) on an identically created memory allocation on the same implementation.

Note

In most cases, it is expected that a non-zero `opaqueAddress` is an address retrieved from [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html) on an identically created memory allocation. If this is not the case, it is likely that `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS` errors will occur.

This is, however, not a strict requirement because trace capture/replay tools may need to adjust memory allocation parameters for imported memory.

If this structure is not present, it is as if `opaqueCaptureAddress` is zero.

Valid Usage (Implicit)

- [](#VUID-VkMemoryOpaqueCaptureAddressAllocateInfo-sType-sType)VUID-VkMemoryOpaqueCaptureAddressAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO`

## [](#_see_also)See Also

[VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryOpaqueCaptureAddressAllocateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0