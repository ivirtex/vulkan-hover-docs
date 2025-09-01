# VkSparseImageOpaqueMemoryBindInfo(3) Manual Page

## Name

VkSparseImageOpaqueMemoryBindInfo - Structure specifying sparse image opaque memory bind information



## [](#_c_specification)C Specification

Memory is bound to opaque regions of `VkImage` objects created with the `VK_IMAGE_CREATE_SPARSE_BINDING_BIT` flag using the following structure:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSparseImageOpaqueMemoryBindInfo {
    VkImage                      image;
    uint32_t                     bindCount;
    const VkSparseMemoryBind*    pBinds;
} VkSparseImageOpaqueMemoryBindInfo;
```

## [](#_members)Members

- `image` is the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object to be bound.
- `bindCount` is the number of [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html) structures in the `pBinds` array.
- `pBinds` is a pointer to an array of [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html) structures.

## [](#_description)Description

Note

This structure is normally used to bind memory to fully-resident sparse images or for mip tail regions of partially resident images. However, it **can** also be used to bind memory for the entire binding range of partially resident images.

If the `pBinds`\[i].flags of an element *i* of `pBinds` does not contain `VK_SPARSE_MEMORY_BIND_METADATA_BIT`, the `resourceOffset` is in the range \[0, [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html)::`size`), This range includes data from all aspects of the image, including metadata. For most implementations this will probably mean that the `resourceOffset` is a simple device address offset within the resource. It is possible for an application to bind a range of memory that includes both resource data and metadata. However, the application would not know what part of the image the memory is used for, or if any range is being used for metadata.

If the `pBinds`\[i].flags of an element *i* of `pBinds` contains `VK_SPARSE_MEMORY_BIND_METADATA_BIT`, the binding range specified **must** be within the mip tail region of the metadata aspect. In this case the `resourceOffset` is not **required** to be a simple device address offset within the resource. However, it *is* defined to be within \[`imageMipTailOffset`, `imageMipTailOffset` + `imageMipTailSize`) for the metadata aspect. See [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html) for the full constraints on binding region with this flag present.

Valid Usage

- [](#VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-01103)VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-01103  
  If the `flags` member of any element of `pBinds` contains `VK_SPARSE_MEMORY_BIND_METADATA_BIT`, the binding range defined **must** be within the mip tail region of the metadata aspect of `image`

Valid Usage (Implicit)

- [](#VUID-VkSparseImageOpaqueMemoryBindInfo-image-parameter)VUID-VkSparseImageOpaqueMemoryBindInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-parameter)VUID-VkSparseImageOpaqueMemoryBindInfo-pBinds-parameter  
  `pBinds` **must** be a valid pointer to an array of `bindCount` valid [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html) structures
- [](#VUID-VkSparseImageOpaqueMemoryBindInfo-bindCount-arraylength)VUID-VkSparseImageOpaqueMemoryBindInfo-bindCount-arraylength  
  `bindCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBindSparseInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindSparseInfo.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkSparseMemoryBind](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseMemoryBind.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseImageOpaqueMemoryBindInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0