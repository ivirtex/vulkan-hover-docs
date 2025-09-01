# VkCuLaunchInfoNVX(3) Manual Page

## Name

VkCuLaunchInfoNVX - Stub description of VkCuLaunchInfoNVX



## [](#_c_specification)C Specification

There is currently no specification language written for this type. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
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

## [](#_members)Members

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkCuLaunchInfoNVX-sType-sType)VUID-VkCuLaunchInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CU_LAUNCH_INFO_NVX`
- [](#VUID-VkCuLaunchInfoNVX-pNext-pNext)VUID-VkCuLaunchInfoNVX-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCuLaunchInfoNVX-function-parameter)VUID-VkCuLaunchInfoNVX-function-parameter  
  `function` **must** be a valid [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html) handle
- [](#VUID-VkCuLaunchInfoNVX-pParams-parameter)VUID-VkCuLaunchInfoNVX-pParams-parameter  
  If `paramCount` is not `0`, `pParams` **must** be a valid pointer to an array of `paramCount` bytes
- [](#VUID-VkCuLaunchInfoNVX-pExtras-parameter)VUID-VkCuLaunchInfoNVX-pExtras-parameter  
  If `extraCount` is not `0`, `pExtras` **must** be a valid pointer to an array of `extraCount` bytes

## [](#_see_also)See Also

[VK\_NVX\_binary\_import](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_binary_import.html), [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCuLaunchKernelNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCuLaunchKernelNVX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCuLaunchInfoNVX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0