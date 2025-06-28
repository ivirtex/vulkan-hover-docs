# VkBufferOpaqueCaptureAddressCreateInfo(3) Manual Page

## Name

VkBufferOpaqueCaptureAddressCreateInfo - Request a specific address for a buffer



## [](#_c_specification)C Specification

To request a specific device address for a buffer, add a [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html) structure to the `pNext` chain of the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structure. The `VkBufferOpaqueCaptureAddressCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkBufferOpaqueCaptureAddressCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           opaqueCaptureAddress;
} VkBufferOpaqueCaptureAddressCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_buffer_device_address
typedef VkBufferOpaqueCaptureAddressCreateInfo VkBufferOpaqueCaptureAddressCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `opaqueCaptureAddress` is the opaque capture address requested for the buffer.

## [](#_description)Description

If `opaqueCaptureAddress` is zero, no specific address is requested.

If `opaqueCaptureAddress` is not zero, then it **should** be an address retrieved from [vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureAddress.html) for an identically created buffer on the same implementation.

If this structure is not present, it is as if `opaqueCaptureAddress` is zero.

Applications **should** avoid creating buffers with application-provided addresses and implementation-provided addresses in the same process, to reduce the likelihood of `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS` errors.

Note

The expected usage for this is that a trace capture/replay tool will add the `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` flag to all buffers that use `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT`, and during capture will save the queried opaque device addresses in the trace. During replay, the buffers will be created specifying the original address so any address values stored in the trace data will remain valid.

Implementations are expected to separate such buffers in the GPU address space so normal allocations will avoid using these addresses. Applications and tools should avoid mixing application-provided and implementation-provided addresses for buffers created with `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`, to avoid address space allocation conflicts.

Valid Usage (Implicit)

- [](#VUID-VkBufferOpaqueCaptureAddressCreateInfo-sType-sType)VUID-VkBufferOpaqueCaptureAddressCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO`

## [](#_see_also)See Also

[VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferOpaqueCaptureAddressCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0