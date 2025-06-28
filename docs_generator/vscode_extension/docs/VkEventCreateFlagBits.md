# VkEventCreateFlagBits(3) Manual Page

## Name

VkEventCreateFlagBits - Event creation flag bits



## [](#_c_specification)C Specification

```c++
// Provided by VK_VERSION_1_0
typedef enum VkEventCreateFlagBits {
  // Provided by VK_VERSION_1_3
    VK_EVENT_CREATE_DEVICE_ONLY_BIT = 0x00000001,
  // Provided by VK_KHR_synchronization2
    VK_EVENT_CREATE_DEVICE_ONLY_BIT_KHR = VK_EVENT_CREATE_DEVICE_ONLY_BIT,
} VkEventCreateFlagBits;
```

## [](#_description)Description

- `VK_EVENT_CREATE_DEVICE_ONLY_BIT` specifies that host event commands will not be used with this event.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkEventCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEventCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkEventCreateFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0