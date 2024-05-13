# VkSemaphoreType(3) Manual Page

## Name

VkSemaphoreType - Specifies the type of a semaphore object



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType`,
specifying the type of a semaphore, are:

``` c
// Provided by VK_VERSION_1_2
typedef enum VkSemaphoreType {
    VK_SEMAPHORE_TYPE_BINARY = 0,
    VK_SEMAPHORE_TYPE_TIMELINE = 1,
  // Provided by VK_KHR_timeline_semaphore
    VK_SEMAPHORE_TYPE_BINARY_KHR = VK_SEMAPHORE_TYPE_BINARY,
  // Provided by VK_KHR_timeline_semaphore
    VK_SEMAPHORE_TYPE_TIMELINE_KHR = VK_SEMAPHORE_TYPE_TIMELINE,
} VkSemaphoreType;
```

or the equivalent

``` c
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreType VkSemaphoreTypeKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SEMAPHORE_TYPE_BINARY` specifies a *binary semaphore* type that
  has a boolean payload indicating whether the semaphore is currently
  signaled or unsignaled. When created, the semaphore is in the
  unsignaled state.

- `VK_SEMAPHORE_TYPE_TIMELINE` specifies a *timeline semaphore* type
  that has a strictly increasing 64-bit unsigned integer payload
  indicating whether the semaphore is signaled with respect to a
  particular reference value. When created, the semaphore payload has
  the value given by the `initialValue` field of
  [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreTypeCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
