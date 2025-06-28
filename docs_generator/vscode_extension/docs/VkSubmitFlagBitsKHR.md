# VkSubmitFlagBits(3) Manual Page

## Name

VkSubmitFlagBits - Bitmask specifying behavior of a submission



## [](#_c_specification)C Specification

Bits which **can** be set in [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html)::`flags`, specifying submission behavior, are:

```c++
// Provided by VK_VERSION_1_3
typedef enum VkSubmitFlagBits {
    VK_SUBMIT_PROTECTED_BIT = 0x00000001,
  // Provided by VK_KHR_synchronization2
    VK_SUBMIT_PROTECTED_BIT_KHR = VK_SUBMIT_PROTECTED_BIT,
} VkSubmitFlagBits;
```

or the equivalent

```c++
// Provided by VK_KHR_synchronization2
typedef VkSubmitFlagBits VkSubmitFlagBitsKHR;
```

## [](#_description)Description

- `VK_SUBMIT_PROTECTED_BIT` specifies that this batch is a protected submission.

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkSubmitFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubmitFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0