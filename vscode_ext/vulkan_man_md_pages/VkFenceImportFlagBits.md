# VkFenceImportFlagBits(3) Manual Page

## Name

VkFenceImportFlagBits - Bitmask specifying additional parameters of
fence payload import



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in

- [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceWin32HandleInfoKHR.html)::`flags`

- [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportFenceFdInfoKHR.html)::`flags`

specifying additional parameters of a fence import operation are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkFenceImportFlagBits {
    VK_FENCE_IMPORT_TEMPORARY_BIT = 0x00000001,
  // Provided by VK_KHR_external_fence
    VK_FENCE_IMPORT_TEMPORARY_BIT_KHR = VK_FENCE_IMPORT_TEMPORARY_BIT,
} VkFenceImportFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_external_fence
typedef VkFenceImportFlagBits VkFenceImportFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_FENCE_IMPORT_TEMPORARY_BIT` specifies that the fence payload will
  be imported only temporarily, as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-fences-importing"
  target="_blank" rel="noopener">Importing Fence Payloads</a>,
  regardless of the permanence of `handleType`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFenceImportFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
