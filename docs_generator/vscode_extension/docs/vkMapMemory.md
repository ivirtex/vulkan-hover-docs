# vkMapMemory(3) Manual Page

## Name

vkMapMemory - Map a memory object into application address space



## [](#_c_specification)C Specification

To retrieve a host virtual address pointer to a region of a mappable memory object, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkMapMemory(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    VkDeviceSize                                offset,
    VkDeviceSize                                size,
    VkMemoryMapFlags                            flags,
    void**                                      ppData);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object to be mapped.
- `offset` is a zero-based byte offset from the beginning of the memory object.
- `size` is the size of the memory range to map, or `VK_WHOLE_SIZE` to map from `offset` to the end of the allocation.
- `flags` is a bitmask of [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlagBits.html) specifying additional parameters of the memory map operation.
- `ppData` is a pointer to a `void*` variable in which a host-accessible pointer to the beginning of the mapped range is returned. The value of the returned pointer minus `offset` **must** be aligned to [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`minMemoryMapAlignment`.

## [](#_description)Description

After a successful call to `vkMapMemory` the memory object `memory` is considered to be currently *host mapped*.

Note

It is an application error to call `vkMapMemory` on a memory object that is already *host mapped*.

Note

`vkMapMemory` will fail if the implementation is unable to allocate an appropriately sized contiguous virtual address range, e.g. due to virtual address space fragmentation or platform limits. In such cases, `vkMapMemory` **must** return `VK_ERROR_MEMORY_MAP_FAILED`. The application **can** improve the likelihood of success by reducing the size of the mapped range and/or removing unneeded mappings using [vkUnmapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUnmapMemory.html).

`vkMapMemory` does not check whether the device memory is currently in use before returning the host-accessible pointer. The application **must** guarantee that any previously submitted command that writes to this range has completed before the host reads from or writes to that range, and that any previously submitted command that reads from that range has completed before the host writes to that region (see [here](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-host-writes) for details on fulfilling such a guarantee). If the device memory was allocated without the `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` set, these guarantees **must** be made for an extended range: the application **must** round down the start of the range to the nearest multiple of [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`, and round the end of the range up to the nearest multiple of [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`nonCoherentAtomSize`.

While a range of device memory is host mapped, the application is responsible for synchronizing both device and host access to that memory range.

Note

It is important for the application developer to become meticulously familiar with all of the mechanisms described in the chapter on [Synchronization and Cache Control](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization) as they are crucial to maintaining memory access ordering.

Calling `vkMapMemory` is equivalent to calling [vkMapMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory2.html) with an empty `pNext` chain.

Valid Usage

- [](#VUID-vkMapMemory-memory-00678)VUID-vkMapMemory-memory-00678  
  `memory` **must** not be currently host mapped
- [](#VUID-vkMapMemory-offset-00679)VUID-vkMapMemory-offset-00679  
  `offset` **must** be less than the size of `memory`
- [](#VUID-vkMapMemory-size-00680)VUID-vkMapMemory-size-00680  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater than `0`
- [](#VUID-vkMapMemory-size-00681)VUID-vkMapMemory-size-00681  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less than or equal to the size of the `memory` minus `offset`
- [](#VUID-vkMapMemory-memory-00682)VUID-vkMapMemory-memory-00682  
  `memory` **must** have been created with a memory type that reports `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT`
- [](#VUID-vkMapMemory-memory-00683)VUID-vkMapMemory-memory-00683  
  `memory` **must** not have been allocated with multiple instances
- [](#VUID-vkMapMemory-flags-09568)VUID-vkMapMemory-flags-09568  
  `VK_MEMORY_MAP_PLACED_BIT_EXT` **must** not be set in `flags`

Valid Usage (Implicit)

- [](#VUID-vkMapMemory-device-parameter)VUID-vkMapMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkMapMemory-memory-parameter)VUID-vkMapMemory-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-vkMapMemory-flags-parameter)VUID-vkMapMemory-flags-parameter  
  `flags` **must** be a valid combination of [VkMemoryMapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlagBits.html) values
- [](#VUID-vkMapMemory-ppData-parameter)VUID-vkMapMemory-ppData-parameter  
  `ppData` **must** be a valid pointer to a pointer value
- [](#VUID-vkMapMemory-memory-parent)VUID-vkMapMemory-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `memory` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_MEMORY_MAP_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkMemoryMapFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryMapFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkMapMemory)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0