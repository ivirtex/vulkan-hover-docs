# VkDirectDriverLoadingInfoLUNARG(3) Manual Page

## Name

VkDirectDriverLoadingInfoLUNARG - Structure specifying the information required to load an additional driver



## [](#_c_specification)C Specification

The `VkDirectDriverLoadingInfoLUNARG` structure is defined as:

```c++
// Provided by VK_LUNARG_direct_driver_loading
typedef struct VkDirectDriverLoadingInfoLUNARG {
    VkStructureType                     sType;
    void*                               pNext;
    VkDirectDriverLoadingFlagsLUNARG    flags;
    PFN_vkGetInstanceProcAddrLUNARG     pfnGetInstanceProcAddr;
} VkDirectDriverLoadingInfoLUNARG;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `pfnGetInstanceProcAddr` is a [PFN\_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkGetInstanceProcAddrLUNARG.html) pointer to the driver [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetInstanceProcAddr.html) function.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDirectDriverLoadingInfoLUNARG-sType-sType)VUID-VkDirectDriverLoadingInfoLUNARG-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_INFO_LUNARG`
- [](#VUID-VkDirectDriverLoadingInfoLUNARG-flags-zerobitmask)VUID-VkDirectDriverLoadingInfoLUNARG-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[PFN\_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkGetInstanceProcAddrLUNARG.html), [VK\_LUNARG\_direct\_driver\_loading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_LUNARG_direct_driver_loading.html), [VkDirectDriverLoadingFlagsLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingFlagsLUNARG.html), [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDirectDriverLoadingInfoLUNARG)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0