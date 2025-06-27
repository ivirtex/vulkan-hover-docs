# VkDeviceQueueCreateFlagBits(3) Manual Page

## Name

VkDeviceQueueCreateFlagBits - Bitmask specifying behavior of the queue



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)::`flags`,
specifying usage behavior of a queue, are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkDeviceQueueCreateFlagBits {
  // Provided by VK_VERSION_1_1
    VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT = 0x00000001,
} VkDeviceQueueCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_QUEUE_CREATE_PROTECTED_BIT` specifies that the device queue
  is a protected-capable queue.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceQueueCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
