# VkTensorTilingARM(3) Manual Page

## Name

VkTensorTilingARM - Specifies the tiling arrangement of data in an tensor



## [](#_c_specification)C Specification

Possible values of [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html)::`tiling`, specifying the tiling arrangement of elements in the tensor, are:

```c++
// Provided by VK_ARM_tensors
typedef enum VkTensorTilingARM {
    VK_TENSOR_TILING_OPTIMAL_ARM = 0,
    VK_TENSOR_TILING_LINEAR_ARM = 1,
} VkTensorTilingARM;
```

## [](#_description)Description

- `VK_TENSOR_TILING_OPTIMAL_ARM` specifies optimal tiling (elements are laid out in an implementation-dependent arrangement, for more efficient memory access).
- `VK_TENSOR_TILING_LINEAR_ARM` specifies linear tiling (elements are laid out linearly and the offset between each element is determined by the [strides](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-tensor-description-strides) of the tensor).

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorTilingARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0