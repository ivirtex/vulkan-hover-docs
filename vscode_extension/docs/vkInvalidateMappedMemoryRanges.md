# vkInvalidateMappedMemoryRanges(3) Manual Page

## Name

vkInvalidateMappedMemoryRanges - Invalidate ranges of mapped memory objects



## [](#_c_specification)C Specification

To invalidate ranges of non-coherent memory from the host caches, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkInvalidateMappedMemoryRanges(
    VkDevice                                    device,
    uint32_t                                    memoryRangeCount,
    const VkMappedMemoryRange*                  pMemoryRanges);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory ranges.
- `memoryRangeCount` is the length of the `pMemoryRanges` array.
- `pMemoryRanges` is a pointer to an array of [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMappedMemoryRange.html) structures describing the memory ranges to invalidate.

## [](#_description)Description

`vkInvalidateMappedMemoryRanges` guarantees that device writes to the memory ranges described by `pMemoryRanges`, which have been made available to the host memory domain using the `VK_ACCESS_HOST_WRITE_BIT` and `VK_ACCESS_HOST_READ_BIT` [access types](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types), are made visible to the host. If a range of non-coherent memory is written by the host and then invalidated without first being flushed, its contents are undefined.

The first [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all host operations that happened-before it, as defined by the host memory model.

Note

This function does not synchronize with device operations directly - other host [synchronization operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization) that depend on device operations such as [vkWaitForFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForFences.html) must be executed beforehand. So for any non-coherent device write to be made visible to the host, there has to be a dependency chain along the following lines:

1. Device write
2. Device memory barrier including host reads in its second scope
3. Signal on the device (e.g. a [fence signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-signaling))
4. Wait on the host (e.g. [vkWaitForFences](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForFences.html))
5. [vkInvalidateMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/latest/man/html/vkInvalidateMappedMemoryRanges.html)

The second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) includes all host operations that happen-after it, as defined by the host memory model.

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is empty.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) includes host reads to the specified memory ranges.

Note

When a device write to a memory location is made visible to the host in this way, each whole aligned set of [`nonCoherentAtomSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nonCoherentAtomSize) bytes that the memory location exists in will also be made visible as if they were written by the device. For example, with a `nonCoherentAtomSize` of 128, if an application writes to the first byte of a memory object on the device, the first 128 bytes of the memory object will be made visible by this command. While the value of the following 127 bytes will be unchanged, this does count as an access for the purpose of synchronization, so care must be taken to avoid data races.

Note

Mapping non-coherent memory does not implicitly invalidate that memory.

Valid Usage (Implicit)

- [](#VUID-vkInvalidateMappedMemoryRanges-device-parameter)VUID-vkInvalidateMappedMemoryRanges-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkInvalidateMappedMemoryRanges-pMemoryRanges-parameter)VUID-vkInvalidateMappedMemoryRanges-pMemoryRanges-parameter  
  `pMemoryRanges` **must** be a valid pointer to an array of `memoryRangeCount` valid [VkMappedMemoryRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMappedMemoryRange.html) structures
- [](#VUID-vkInvalidateMappedMemoryRanges-memoryRangeCount-arraylength)VUID-vkInvalidateMappedMemoryRanges-memoryRangeCount-arraylength  
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

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkInvalidateMappedMemoryRanges)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0