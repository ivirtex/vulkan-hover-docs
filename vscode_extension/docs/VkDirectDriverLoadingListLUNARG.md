# VkDirectDriverLoadingListLUNARG(3) Manual Page

## Name

VkDirectDriverLoadingListLUNARG - Structure specifying additional drivers to load



## [](#_c_specification)C Specification

The `VkDirectDriverLoadingListLUNARG` structure is defined as:

```c++
// Provided by VK_LUNARG_direct_driver_loading
typedef struct VkDirectDriverLoadingListLUNARG {
    VkStructureType                           sType;
    const void*                               pNext;
    VkDirectDriverLoadingModeLUNARG           mode;
    uint32_t                                  driverCount;
    const VkDirectDriverLoadingInfoLUNARG*    pDrivers;
} VkDirectDriverLoadingListLUNARG;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `mode` controls the mode in which to load the provided drivers.
- `driverCount` is the number of driver manifest paths.
- `pDrivers` is a pointer to an array of `driverCount` [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingInfoLUNARG.html) structures.

## [](#_description)Description

When creating a Vulkan instance for which additional drivers are to be included, add a `VkDirectDriverLoadingListLUNARG` structure to the pNext chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure, and include in it the list of `VkDirectDriverLoadingInfoLUNARG` structures which contain the information necessary to load additional drivers.

Valid Usage (Implicit)

- [](#VUID-VkDirectDriverLoadingListLUNARG-sType-sType)VUID-VkDirectDriverLoadingListLUNARG-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DIRECT_DRIVER_LOADING_LIST_LUNARG`
- [](#VUID-VkDirectDriverLoadingListLUNARG-mode-parameter)VUID-VkDirectDriverLoadingListLUNARG-mode-parameter  
  `mode` **must** be a valid [VkDirectDriverLoadingModeLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingModeLUNARG.html) value
- [](#VUID-VkDirectDriverLoadingListLUNARG-pDrivers-parameter)VUID-VkDirectDriverLoadingListLUNARG-pDrivers-parameter  
  `pDrivers` **must** be a valid pointer to an array of `driverCount` valid [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingInfoLUNARG.html) structures
- [](#VUID-VkDirectDriverLoadingListLUNARG-driverCount-arraylength)VUID-VkDirectDriverLoadingListLUNARG-driverCount-arraylength  
  `driverCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_LUNARG\_direct\_driver\_loading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_LUNARG_direct_driver_loading.html), [VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingInfoLUNARG.html), [VkDirectDriverLoadingModeLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingModeLUNARG.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDirectDriverLoadingListLUNARG).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0