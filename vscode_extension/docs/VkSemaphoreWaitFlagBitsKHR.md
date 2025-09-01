# VkSemaphoreWaitFlagBits(3) Manual Page

## Name

VkSemaphoreWaitFlagBits - Bitmask specifying additional parameters of a semaphore wait operation



## [](#_c_specification)C Specification

Bits which **can** be set in [VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitInfo.html)::`flags`, specifying additional parameters of a semaphore wait operation, are:

```c++
// Provided by VK_VERSION_1_2
typedef enum VkSemaphoreWaitFlagBits {
    VK_SEMAPHORE_WAIT_ANY_BIT = 0x00000001,
  // Provided by VK_KHR_timeline_semaphore
    VK_SEMAPHORE_WAIT_ANY_BIT_KHR = VK_SEMAPHORE_WAIT_ANY_BIT,
} VkSemaphoreWaitFlagBits;
```

or the equivalent

```c++
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreWaitFlagBits VkSemaphoreWaitFlagBitsKHR;
```

## [](#_description)Description

- `VK_SEMAPHORE_WAIT_ANY_BIT` specifies that the semaphore wait condition is that at least one of the semaphores in `VkSemaphoreWaitInfo`::`pSemaphores` has reached the value specified by the corresponding element of `VkSemaphoreWaitInfo`::`pValues`. If `VK_SEMAPHORE_WAIT_ANY_BIT` is not set, the semaphore wait condition is that all of the semaphores in `VkSemaphoreWaitInfo`::`pSemaphores` have reached the value specified by the corresponding element of `VkSemaphoreWaitInfo`::`pValues`.

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkSemaphoreWaitFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreWaitFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreWaitFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0