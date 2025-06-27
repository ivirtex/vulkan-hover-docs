# VkSemaphoreImportFlagBits(3) Manual Page

## Name

VkSemaphoreImportFlagBits - Bitmask specifying additional parameters of
semaphore payload import



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in

- [VkImportSemaphoreWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreWin32HandleInfoKHR.html)::`flags`

- [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreFdInfoKHR.html)::`flags`

- [VkImportSemaphoreZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportSemaphoreZirconHandleInfoFUCHSIA.html)::`flags`

specifying additional parameters of a semaphore import operation are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkSemaphoreImportFlagBits {
    VK_SEMAPHORE_IMPORT_TEMPORARY_BIT = 0x00000001,
  // Provided by VK_KHR_external_semaphore
    VK_SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR = VK_SEMAPHORE_IMPORT_TEMPORARY_BIT,
} VkSemaphoreImportFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_external_semaphore
typedef VkSemaphoreImportFlagBits VkSemaphoreImportFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

These bits have the following meanings:

- `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT` specifies that the semaphore
  payload will be imported only temporarily, as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-importing"
  target="_blank" rel="noopener">Importing Semaphore Payloads</a>,
  regardless of the permanence of `handleType`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreImportFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
