# VkDirectDriverLoadingListLUNARG(3) Manual Page

## Name

VkDirectDriverLoadingListLUNARG - Structure specifying additional
drivers to load



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDirectDriverLoadingListLUNARG` structure is defined as:

``` c
// Provided by VK_LUNARG_direct_driver_loading
typedef struct VkDirectDriverLoadingListLUNARG {
    VkStructureType                           sType;
    const void*                               pNext;
    VkDirectDriverLoadingModeLUNARG           mode;
    uint32_t                                  driverCount;
    const VkDirectDriverLoadingInfoLUNARG*    pDrivers;
} VkDirectDriverLoadingListLUNARG;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `mode` controls the mode in which to load the provided drivers.

- `driverCount` is the number of driver manifest paths.

- `pDrivers` is a pointer to an array of `driverCount`
  [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingInfoLUNARG.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

When creating a Vulkan instance for which additional drivers are to be
included, add a `VkDirectDriverLoadingListLUNARG` structure to the pNext
chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html)
structure, and include in it the list of
`VkDirectDriverLoadingInfoLUNARG` structures which contain the
information necessary to load additional drivers.

Valid Usage (Implicit)

- <a href="#VUID-VkDirectDriverLoadingListLUNARG-sType-sType"
  id="VUID-VkDirectDriverLoadingListLUNARG-sType-sType"></a>
  VUID-VkDirectDriverLoadingListLUNARG-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_LIST_LUNARG`

- <a href="#VUID-VkDirectDriverLoadingListLUNARG-mode-parameter"
  id="VUID-VkDirectDriverLoadingListLUNARG-mode-parameter"></a>
  VUID-VkDirectDriverLoadingListLUNARG-mode-parameter  
  `mode` **must** be a valid
  [VkDirectDriverLoadingModeLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingModeLUNARG.html)
  value

- <a href="#VUID-VkDirectDriverLoadingListLUNARG-pDrivers-parameter"
  id="VUID-VkDirectDriverLoadingListLUNARG-pDrivers-parameter"></a>
  VUID-VkDirectDriverLoadingListLUNARG-pDrivers-parameter  
  `pDrivers` **must** be a valid pointer to an array of `driverCount`
  valid
  [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingInfoLUNARG.html)
  structures

- <a href="#VUID-VkDirectDriverLoadingListLUNARG-driverCount-arraylength"
  id="VUID-VkDirectDriverLoadingListLUNARG-driverCount-arraylength"></a>
  VUID-VkDirectDriverLoadingListLUNARG-driverCount-arraylength  
  `driverCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_LUNARG_direct_driver_loading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUNARG_direct_driver_loading.html),
[VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingInfoLUNARG.html),
[VkDirectDriverLoadingModeLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingModeLUNARG.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDirectDriverLoadingListLUNARG"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
