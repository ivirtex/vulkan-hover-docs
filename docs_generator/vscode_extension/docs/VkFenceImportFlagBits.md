# VkFenceImportFlagBits(3) Manual Page

## Name

VkFenceImportFlagBits - Bitmask specifying additional parameters of fence payload import



## [](#_c_specification)C Specification

Bits which **can** be set in

- [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceWin32HandleInfoKHR.html)::`flags`
- [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceFdInfoKHR.html)::`flags`

specifying additional parameters of a fence import operation are:

```c++
// Provided by VK_VERSION_1_1
typedef enum VkFenceImportFlagBits {
    VK_FENCE_IMPORT_TEMPORARY_BIT = 0x00000001,
  // Provided by VK_KHR_external_fence
    VK_FENCE_IMPORT_TEMPORARY_BIT_KHR = VK_FENCE_IMPORT_TEMPORARY_BIT,
} VkFenceImportFlagBits;
```

or the equivalent

```c++
// Provided by VK_KHR_external_fence
typedef VkFenceImportFlagBits VkFenceImportFlagBitsKHR;
```

## [](#_description)Description

- `VK_FENCE_IMPORT_TEMPORARY_BIT` specifies that the fence payload will be imported only temporarily, as described in [Importing Fence Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-fences-importing), regardless of the permanence of `handleType`.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceImportFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFenceImportFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0