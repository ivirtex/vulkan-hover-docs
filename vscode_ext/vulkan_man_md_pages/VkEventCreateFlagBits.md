# VkEventCreateFlagBits(3) Manual Page

## Name

VkEventCreateFlagBits - Event creation flag bits



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
// Provided by VK_VERSION_1_0
typedef enum VkEventCreateFlagBits {
  // Provided by VK_VERSION_1_3
    VK_EVENT_CREATE_DEVICE_ONLY_BIT = 0x00000001,
  // Provided by VK_KHR_synchronization2
    VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR = VK_EVENT_CREATE_DEVICE_ONLY_BIT,
} VkEventCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_EVENT_CREATE_DEVICE_ONLY_BIT` specifies that host event commands
  will not be used with this event.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkEventCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkEventCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
