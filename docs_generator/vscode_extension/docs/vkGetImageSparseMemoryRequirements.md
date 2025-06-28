# vkGetImageSparseMemoryRequirements(3) Manual Page

## Name

vkGetImageSparseMemoryRequirements - Query the memory requirements for a sparse image



## [](#_c_specification)C Specification

To query sparse memory requirements for an image, call:

```c++
// Provided by VK_VERSION_1_0
void vkGetImageSparseMemoryRequirements(
    VkDevice                                    device,
    VkImage                                     image,
    uint32_t*                                   pSparseMemoryRequirementCount,
    VkSparseImageMemoryRequirements*            pSparseMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image.
- `image` is the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object to get the memory requirements for.
- `pSparseMemoryRequirementCount` is a pointer to an integer related to the number of sparse memory requirements available or queried, as described below.
- `pSparseMemoryRequirements` is either `NULL` or a pointer to an array of `VkSparseImageMemoryRequirements` structures.

## [](#_description)Description

If `pSparseMemoryRequirements` is `NULL`, then the number of sparse memory requirements available is returned in `pSparseMemoryRequirementCount`. Otherwise, `pSparseMemoryRequirementCount` **must** point to a variable set by the application to the number of elements in the `pSparseMemoryRequirements` array, and on return the variable is overwritten with the number of structures actually written to `pSparseMemoryRequirements`. If `pSparseMemoryRequirementCount` is less than the number of sparse memory requirements available, at most `pSparseMemoryRequirementCount` structures will be written.

If the image was not created with `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` then `pSparseMemoryRequirementCount` will be zero and `pSparseMemoryRequirements` will not be written to.

Note

It is legal for an implementation to report a larger value in `VkMemoryRequirements`::`size` than would be obtained by adding together memory sizes for all `VkSparseImageMemoryRequirements` returned by `vkGetImageSparseMemoryRequirements`. This **may** occur when the implementation requires unused padding in the address range describing the resource.

Valid Usage (Implicit)

- [](#VUID-vkGetImageSparseMemoryRequirements-device-parameter)VUID-vkGetImageSparseMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetImageSparseMemoryRequirements-image-parameter)VUID-vkGetImageSparseMemoryRequirements-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter)VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirementCount-parameter  
  `pSparseMemoryRequirementCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter)VUID-vkGetImageSparseMemoryRequirements-pSparseMemoryRequirements-parameter  
  If the value referenced by `pSparseMemoryRequirementCount` is not `0`, and `pSparseMemoryRequirements` is not `NULL`, `pSparseMemoryRequirements` **must** be a valid pointer to an array of `pSparseMemoryRequirementCount` [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements.html) structures
- [](#VUID-vkGetImageSparseMemoryRequirements-image-parent)VUID-vkGetImageSparseMemoryRequirements-image-parent  
  `image` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetImageSparseMemoryRequirements)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0