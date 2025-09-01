# VkDirectDriverLoadingModeLUNARG(3) Manual Page

## Name

VkDirectDriverLoadingModeLUNARG - Specify loader behavior of added drivers



## [](#_c_specification)C Specification

Possible values of [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html)::`mode`, specifying the mode in which drivers are used, are:

```c++
// Provided by VK_LUNARG_direct_driver_loading
typedef enum VkDirectDriverLoadingModeLUNARG {
    VK_DIRECT_DRIVER_LOADING_MODE_EXCLUSIVE_LUNARG = 0,
    VK_DIRECT_DRIVER_LOADING_MODE_INCLUSIVE_LUNARG = 1,
} VkDirectDriverLoadingModeLUNARG;
```

## [](#_description)Description

- `VK_DIRECT_DRIVER_LOADING_MODE_EXCLUSIVE_LUNARG` specifies that the provided drivers are used instead of the system-loaded drivers.
- `VK_DIRECT_DRIVER_LOADING_MODE_INCLUSIVE_LUNARG` specifies that the provided drivers are used in addition to the system-loaded drivers.

## [](#_see_also)See Also

[VK\_LUNARG\_direct\_driver\_loading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_LUNARG_direct_driver_loading.html), [VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDirectDriverLoadingListLUNARG.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDirectDriverLoadingModeLUNARG).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0