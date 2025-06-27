# VkDirectDriverLoadingModeLUNARG(3) Manual Page

## Name

VkDirectDriverLoadingModeLUNARG - Specify loader behavior of added
drivers



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html)::`mode`,
specifying the mode in which drivers are used, are:

``` c
// Provided by VK_LUNARG_direct_driver_loading
typedef enum VkDirectDriverLoadingModeLUNARG {
    VK_DIRECT_DRIVER_LOADING_MODE_EXCLUSIVE_LUNARG = 0,
    VK_DIRECT_DRIVER_LOADING_MODE_INCLUSIVE_LUNARG = 1,
} VkDirectDriverLoadingModeLUNARG;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DIRECT_DRIVER_LOADING_MODE_EXCLUSIVE_LUNARG` specifies that the
  provided drivers are used instead of the system-loaded drivers.

- `VK_DIRECT_DRIVER_LOADING_MODE_INCLUSIVE_LUNARG` specifies that the
  provided drivers are used in addition to the system-loaded drivers.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_LUNARG_direct_driver_loading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUNARG_direct_driver_loading.html),
[VkDirectDriverLoadingListLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingListLUNARG.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDirectDriverLoadingModeLUNARG"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
