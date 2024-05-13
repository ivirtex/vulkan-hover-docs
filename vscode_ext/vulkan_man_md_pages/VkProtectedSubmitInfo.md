# VkProtectedSubmitInfo(3) Manual Page

## Name

VkProtectedSubmitInfo - Structure indicating whether the submission is
protected



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html) includes a
`VkProtectedSubmitInfo` structure, then the structure indicates whether
the batch is protected. The `VkProtectedSubmitInfo` structure is defined
as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkProtectedSubmitInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           protectedSubmit;
} VkProtectedSubmitInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `protectedSubmit` specifies whether the batch is protected. If
  `protectedSubmit` is `VK_TRUE`, the batch is protected. If
  `protectedSubmit` is `VK_FALSE`, the batch is unprotected. If the
  `VkSubmitInfo`::`pNext` chain does not include this structure, the
  batch is unprotected.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkProtectedSubmitInfo-sType-sType"
  id="VUID-VkProtectedSubmitInfo-sType-sType"></a>
  VUID-VkProtectedSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PROTECTED_SUBMIT_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkProtectedSubmitInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
