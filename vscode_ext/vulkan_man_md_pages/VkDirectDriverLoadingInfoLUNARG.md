# VkDirectDriverLoadingInfoLUNARG(3) Manual Page

## Name

VkDirectDriverLoadingInfoLUNARG - Structure specifying the information
required to load an additional driver



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDirectDriverLoadingInfoLUNARG` structure is defined as:

``` c
// Provided by VK_LUNARG_direct_driver_loading
typedef struct VkDirectDriverLoadingInfoLUNARG {
    VkStructureType                     sType;
    void*                               pNext;
    VkDirectDriverLoadingFlagsLUNARG    flags;
    PFN_vkGetInstanceProcAddrLUNARG     pfnGetInstanceProcAddr;
} VkDirectDriverLoadingInfoLUNARG;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `pfnGetInstanceProcAddr` is a
  [PFN_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkGetInstanceProcAddrLUNARG.html)
  pointer to the driver
  [vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetInstanceProcAddr.html) function.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDirectDriverLoadingInfoLUNARG-sType-sType"
  id="VUID-VkDirectDriverLoadingInfoLUNARG-sType-sType"></a>
  VUID-VkDirectDriverLoadingInfoLUNARG-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_INFO_LUNARG`

- <a href="#VUID-VkDirectDriverLoadingInfoLUNARG-flags-zerobitmask"
  id="VUID-VkDirectDriverLoadingInfoLUNARG-flags-zerobitmask"></a>
  VUID-VkDirectDriverLoadingInfoLUNARG-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkGetInstanceProcAddrLUNARG.html),
[VK_LUNARG_direct_driver_loading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUNARG_direct_driver_loading.html),
[VkDirectDriverLoadingFlagsLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingFlagsLUNARG.html),
[VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDirectDriverLoadingInfoLUNARG"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
