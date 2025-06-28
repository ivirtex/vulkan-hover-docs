# VkBufferMemoryBarrier(3) Manual Page

## Name

VkBufferMemoryBarrier - Structure specifying a buffer memory barrier



## [](#_c_specification)C Specification

The `VkBufferMemoryBarrier` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkBufferMemoryBarrier {
    VkStructureType    sType;
    const void*        pNext;
    VkAccessFlags      srcAccessMask;
    VkAccessFlags      dstAccessMask;
    uint32_t           srcQueueFamilyIndex;
    uint32_t           dstQueueFamilyIndex;
    VkBuffer           buffer;
    VkDeviceSize       offset;
    VkDeviceSize       size;
} VkBufferMemoryBarrier;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `dstAccessMask` is a bitmask of [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits.html) specifying a [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks).
- `srcQueueFamilyIndex` is the source queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `dstQueueFamilyIndex` is the destination queue family for a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).
- `buffer` is a handle to the buffer whose backing memory is affected by the barrier.
- `offset` is an offset in bytes into the backing memory for `buffer`; this is relative to the base offset as bound to the buffer (see [vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory.html)).
- `size` is a size in bytes of the affected area of backing memory for `buffer`, or `VK_WHOLE_SIZE` to use the range from `offset` to the end of the buffer.

## [](#_description)Description

The first [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to access to memory through the specified buffer range, via access types in the [source access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `srcAccessMask` and, if a [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html) is passed in `pNext`, `srcAccessMask3`. If the source access mask includes `VK_ACCESS_HOST_WRITE_BIT`, a [memory domain operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-available-and-visible) is performed where available memory in the host domain is also made available to the device domain.

The second [access scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-access-scopes) is limited to access to memory through the specified buffer range, via access types in the [destination access mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-masks) specified by `dstAccessMask` and, if a [VkMemoryBarrierAccessFlags3KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryBarrierAccessFlags3KHR.html) is passed in `pNext`, `dstAccessMask3`. If the destination access mask includes `VK_ACCESS_HOST_WRITE_BIT` or `VK_ACCESS_HOST_READ_BIT`, a [memory domain operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-available-and-visible) is performed where available memory in the device domain is also made available to the host domain.

Note

When `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` is used, available memory in host domain is automatically made visible to host domain, and any host write is automatically made available to host domain.

If `srcQueueFamilyIndex` is not equal to `dstQueueFamilyIndex`, and `srcQueueFamilyIndex` is equal to the current queue family, then the memory barrier defines a [queue family release operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-release) for the specified buffer range, and if `dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the second synchronization scope of the calling command does not apply to this operation.

If `dstQueueFamilyIndex` is not equal to `srcQueueFamilyIndex`, and `dstQueueFamilyIndex` is equal to the current queue family, then the memory barrier defines a [queue family acquire operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers-acquire) for the specified buffer range, and if `dependencyFlags` did not include `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR`, the first synchronization scope of the calling command does not apply to this operation.

Valid Usage

- [](#VUID-VkBufferMemoryBarrier-offset-01187)VUID-VkBufferMemoryBarrier-offset-01187  
  `offset` **must** be less than the size of `buffer`
- [](#VUID-VkBufferMemoryBarrier-size-01188)VUID-VkBufferMemoryBarrier-size-01188  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be greater than `0`
- [](#VUID-VkBufferMemoryBarrier-size-01189)VUID-VkBufferMemoryBarrier-size-01189  
  If `size` is not equal to `VK_WHOLE_SIZE`, `size` **must** be less than or equal to than the size of `buffer` minus `offset`
- [](#VUID-VkBufferMemoryBarrier-buffer-01931)VUID-VkBufferMemoryBarrier-buffer-01931  
  If `buffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkBufferMemoryBarrier-buffer-09095)VUID-VkBufferMemoryBarrier-buffer-09095  
  If `buffer` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `srcQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkBufferMemoryBarrier-buffer-09096)VUID-VkBufferMemoryBarrier-buffer-09096  
  If `buffer` was created with a sharing mode of `VK_SHARING_MODE_EXCLUSIVE`, and `srcQueueFamilyIndex` and `dstQueueFamilyIndex` are not equal, `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_EXTERNAL`, `VK_QUEUE_FAMILY_FOREIGN_EXT`, or a valid queue family
- [](#VUID-VkBufferMemoryBarrier-None-09097)VUID-VkBufferMemoryBarrier-None-09097  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkBufferMemoryBarrier-None-09098)VUID-VkBufferMemoryBarrier-None-09098  
  If the [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html) extension is not enabled, and the value of [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkApplicationInfo.html)::`apiVersion` used to create the [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) is not greater than or equal to Version 1.1, `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-09099)VUID-VkBufferMemoryBarrier-srcQueueFamilyIndex-09099  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `srcQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkBufferMemoryBarrier-dstQueueFamilyIndex-09100)VUID-VkBufferMemoryBarrier-dstQueueFamilyIndex-09100  
  If the [VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html) extension is not enabled `dstQueueFamilyIndex` **must** not be `VK_QUEUE_FAMILY_FOREIGN_EXT`
- [](#VUID-VkBufferMemoryBarrier-None-09049)VUID-VkBufferMemoryBarrier-None-09049  
  If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is not enabled, and `buffer` was created with a sharing mode of `VK_SHARING_MODE_CONCURRENT`, at least one of `srcQueueFamilyIndex` and `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED`
- [](#VUID-VkBufferMemoryBarrier-None-09050)VUID-VkBufferMemoryBarrier-None-09050  
  If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is not enabled, and `buffer` was created with a sharing mode of `VK_SHARING_MODE_CONCURRENT`, `srcQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`
- [](#VUID-VkBufferMemoryBarrier-None-09051)VUID-VkBufferMemoryBarrier-None-09051  
  If the [`synchronization2`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-synchronization2) feature is not enabled, and `buffer` was created with a sharing mode of `VK_SHARING_MODE_CONCURRENT`, `dstQueueFamilyIndex` **must** be `VK_QUEUE_FAMILY_IGNORED` or `VK_QUEUE_FAMILY_EXTERNAL`

Valid Usage (Implicit)

- [](#VUID-VkBufferMemoryBarrier-sType-sType)VUID-VkBufferMemoryBarrier-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_MEMORY_BARRIER`
- [](#VUID-VkBufferMemoryBarrier-pNext-pNext)VUID-VkBufferMemoryBarrier-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkExternalMemoryAcquireUnmodifiedEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryAcquireUnmodifiedEXT.html)
- [](#VUID-VkBufferMemoryBarrier-sType-unique)VUID-VkBufferMemoryBarrier-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBufferMemoryBarrier-buffer-parameter)VUID-VkBufferMemoryBarrier-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAccessFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlags.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPipelineBarrier.html), [vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWaitEvents.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferMemoryBarrier)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0