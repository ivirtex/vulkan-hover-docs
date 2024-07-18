# VkSemaphoreWaitFlagBits(3) Manual Page

## Name

VkSemaphoreWaitFlagBits - Bitmask specifying additional parameters of a
semaphore wait operation



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkSemaphoreWaitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitInfo.html)::`flags`, specifying
additional parameters of a semaphore wait operation, are:

``` c
// Provided by VK_VERSION_1_2
typedef enum VkSemaphoreWaitFlagBits {
    VK_SEMAPHORE_WAIT_ANY_BIT = 0x00000001,
  // Provided by VK_KHR_timeline_semaphore
    VK_SEMAPHORE_WAIT_ANY_BIT_KHR = VK_SEMAPHORE_WAIT_ANY_BIT,
} VkSemaphoreWaitFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreWaitFlagBits VkSemaphoreWaitFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SEMAPHORE_WAIT_ANY_BIT` specifies that the semaphore wait
  condition is that at least one of the semaphores in
  `VkSemaphoreWaitInfo`::`pSemaphores` has reached the value specified
  by the corresponding element of `VkSemaphoreWaitInfo`::`pValues`. If
  `VK_SEMAPHORE_WAIT_ANY_BIT` is not set, the semaphore wait condition
  is that all of the semaphores in `VkSemaphoreWaitInfo`::`pSemaphores`
  have reached the value specified by the corresponding element of
  `VkSemaphoreWaitInfo`::`pValues`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkSemaphoreWaitFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreWaitFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
