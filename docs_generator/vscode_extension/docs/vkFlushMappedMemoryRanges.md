# vkFlushMappedMemoryRanges(3) Manual Page

## Name

vkFlushMappedMemoryRanges - Flush mapped memory ranges



## [](#_c_specification)C Specification

To flush ranges of non-coherent memory from the host caches, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkFlushMappedMemoryRanges(
    VkDevice                                    device,
    uint32_t                                    memoryRangeCount,
    const VkMappedMemoryRange*                  pMemoryRanges);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory ranges.
- `memoryRangeCount` is the length of the `pMemoryRanges` array.
- `pMemoryRanges` is a pointer to an array of [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMappedMemoryRange.html) structures describing the memory ranges to flush.

## [](#_description)Description

`vkFlushMappedMemoryRanges` guarantees that host writes to the memory ranges described by `pMemoryRanges` are made available to the host memory domain, such that they **can** be made available to the device memory domain via [memory domain operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-available-and-visible) using the `VK_ACCESS_HOST_WRITE_BIT` [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types).

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all host operations that happened-before it, as defined by the host memory model.

Note

Some systems allow writes that do not directly integrate with the host memory model; these have to be synchronized by the application manually. One example of this is non-temporal store instructions on x86; to ensure these happen-before submission, applications should call `_mm_sfence()`.

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) is empty.

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) includes host writes to the specified memory ranges.

Note

When a host write to a memory location is made available in this way, each whole aligned set of [`nonCoherentAtomSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nonCoherentAtomSize) bytes that the memory location exists in will also be made available as if they were written by the host. For example, with a `nonCoherentAtomSize` of 128, if an application writes to the first byte of a memory object via a host mapping, the first 128 bytes of the memory object will be made available by this command. While the value of the following 127 bytes will be unchanged, this does count as an access for the purpose of synchronization, so care must be taken to avoid data races.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is empty.

Unmapping non-coherent memory does not implicitly flush the host mapped memory, and host writes that have not been flushed **may** not ever be visible to the device. However, implementations **must** ensure that writes that have not been flushed do not become visible to any other memory.

Note

The above guarantee avoids a potential memory corruption in scenarios where host writes to a mapped memory object have not been flushed before the memory is unmapped (or freed), and the virtual address range is subsequently reused for a different mapping (or memory allocation).

Valid Usage (Implicit)

- [](#VUID-vkFlushMappedMemoryRanges-device-parameter)VUID-vkFlushMappedMemoryRanges-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkFlushMappedMemoryRanges-pMemoryRanges-parameter)VUID-vkFlushMappedMemoryRanges-pMemoryRanges-parameter  
  `pMemoryRanges` **must** be a valid pointer to an array of `memoryRangeCount` valid [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMappedMemoryRange.html) structures
- [](#VUID-vkFlushMappedMemoryRanges-memoryRangeCount-arraylength)VUID-vkFlushMappedMemoryRanges-memoryRangeCount-arraylength  
  `memoryRangeCount` **must** be greater than `0`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMappedMemoryRange.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkFlushMappedMemoryRanges)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0