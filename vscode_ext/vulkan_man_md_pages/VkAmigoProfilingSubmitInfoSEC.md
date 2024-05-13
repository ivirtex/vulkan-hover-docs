# VkAmigoProfilingSubmitInfoSEC(3) Manual Page

## Name

VkAmigoProfilingSubmitInfoSEC - Stub description of
VkAmigoProfilingSubmitInfoSEC



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this type. This
section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_SEC_amigo_profiling
typedef struct VkAmigoProfilingSubmitInfoSEC {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           firstDrawTimestamp;
    uint64_t           swapBufferTimestamp;
} VkAmigoProfilingSubmitInfoSEC;
```

## <a href="#_members" class="anchor"></a>Members

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkAmigoProfilingSubmitInfoSEC-sType-sType"
  id="VUID-VkAmigoProfilingSubmitInfoSEC-sType-sType"></a>
  VUID-VkAmigoProfilingSubmitInfoSEC-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_AMIGO_PROFILING_SUBMIT_INFO_SEC`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_SEC_amigo_profiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_SEC_amigo_profiling.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAmigoProfilingSubmitInfoSEC"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
