# VkCuLaunchInfoNVX(3) Manual Page

## Name

VkCuLaunchInfoNVX - Stub description of VkCuLaunchInfoNVX



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this type. This
section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_NVX_binary_import
typedef struct VkCuLaunchInfoNVX {
    VkStructureType        sType;
    const void*            pNext;
    VkCuFunctionNVX        function;
    uint32_t               gridDimX;
    uint32_t               gridDimY;
    uint32_t               gridDimZ;
    uint32_t               blockDimX;
    uint32_t               blockDimY;
    uint32_t               blockDimZ;
    uint32_t               sharedMemBytes;
    size_t                 paramCount;
    const void* const *    pParams;
    size_t                 extraCount;
    const void* const *    pExtras;
} VkCuLaunchInfoNVX;
```

## <a href="#_members" class="anchor"></a>Members

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkCuLaunchInfoNVX-sType-sType"
  id="VUID-VkCuLaunchInfoNVX-sType-sType"></a>
  VUID-VkCuLaunchInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CU_LAUNCH_INFO_NVX`

- <a href="#VUID-VkCuLaunchInfoNVX-pNext-pNext"
  id="VUID-VkCuLaunchInfoNVX-pNext-pNext"></a>
  VUID-VkCuLaunchInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCuLaunchInfoNVX-function-parameter"
  id="VUID-VkCuLaunchInfoNVX-function-parameter"></a>
  VUID-VkCuLaunchInfoNVX-function-parameter  
  `function` **must** be a valid [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuFunctionNVX.html)
  handle

- <a href="#VUID-VkCuLaunchInfoNVX-pParams-parameter"
  id="VUID-VkCuLaunchInfoNVX-pParams-parameter"></a>
  VUID-VkCuLaunchInfoNVX-pParams-parameter  
  If `paramCount` is not `0`, `pParams` **must** be a valid pointer to
  an array of `paramCount` bytes

- <a href="#VUID-VkCuLaunchInfoNVX-pExtras-parameter"
  id="VUID-VkCuLaunchInfoNVX-pExtras-parameter"></a>
  VUID-VkCuLaunchInfoNVX-pExtras-parameter  
  If `extraCount` is not `0`, `pExtras` **must** be a valid pointer to
  an array of `extraCount` bytes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NVX_binary_import](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_binary_import.html),
[VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuFunctionNVX.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCuLaunchKernelNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCuLaunchKernelNVX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCuLaunchInfoNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
