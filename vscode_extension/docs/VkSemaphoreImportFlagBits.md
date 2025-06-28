# VkSemaphoreImportFlagBits(3) Manual Page

## Name

VkSemaphoreImportFlagBits - Bitmask specifying additional parameters of semaphore payload import



## [](#_c_specification)C Specification

Bits which **can** be set in

- [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)::`flags`
- [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreFdInfoKHR.html)::`flags`
- [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)::`flags`

specifying additional parameters of a semaphore import operation are:

```c++
// Provided by VK_VERSION_1_1
typedef enum VkSemaphoreImportFlagBits {
    VK_SEMAPHORE_IMPORT_TEMPORARY_BIT = 0x00000001,
  // Provided by VK_KHR_external_semaphore
    VK_SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR = VK_SEMAPHORE_IMPORT_TEMPORARY_BIT,
} VkSemaphoreImportFlagBits;
```

or the equivalent

```c++
// Provided by VK_KHR_external_semaphore
typedef VkSemaphoreImportFlagBits VkSemaphoreImportFlagBitsKHR;
```

## [](#_description)Description

These bits have the following meanings:

- `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT` specifies that the semaphore payload will be imported only temporarily, as described in [Importing Semaphore Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing), regardless of the permanence of `handleType`.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreImportFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0