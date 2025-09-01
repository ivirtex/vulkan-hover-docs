# VkTensorViewCreateInfoARM(3) Manual Page

## Name

VkTensorViewCreateInfoARM - Structure specifying parameters of a newly created tensor view



## [](#_c_specification)C Specification

The `VkTensorViewCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorViewCreateInfoARM {
    VkStructureType               sType;
    const void*                   pNext;
    VkTensorViewCreateFlagsARM    flags;
    VkTensorARM                   tensor;
    VkFormat                      format;
} VkTensorViewCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `tensor` is a [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) on which the view will be created.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) describing the format and type used to interpret elements in the tensor.

## [](#_description)Description

If `tensor` was created with the `VK_TENSOR_CREATE_MUTABLE_FORMAT_BIT_ARM` flag, `format` **can** be different from the tensor’s format, but if they are not equal they **must** be *compatible*. Tensor format compatibility is defined in the [Format Compatibility Classes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes) section. Views of compatible formats will have the same mapping between element locations irrespective of the `format`, with only the interpretation of the bit pattern changing.

Note

Values intended to be used with one view format **may** not be exactly preserved when written or read through a different format. For example, an integer value that happens to have the bit pattern of a floating-point denorm or NaN **may** be flushed or canonicalized when written or read through a view with a floating-point format. Similarly, a value written through a signed normalized format that has a bit pattern exactly equal to -2b **may** be changed to -2b + 1 as described in [Conversion from Normalized Fixed-Point to Floating-Point](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-fixedfpconv).

Valid Usage

- [](#VUID-VkTensorViewCreateInfoARM-tensor-09743)VUID-VkTensorViewCreateInfoARM-tensor-09743  
  If `tensor` was not created with `VK_TENSOR_CREATE_MUTABLE_FORMAT_BIT_ARM` flag, `format` **must** be identical to the `format` used to create `tensor`
- [](#VUID-VkTensorViewCreateInfoARM-tensor-09744)VUID-VkTensorViewCreateInfoARM-tensor-09744  
  If `tensor` was created with `VK_TENSOR_CREATE_MUTABLE_FORMAT_BIT_ARM` flag, `format` **must** be compatible with the `format` used to create `tensor`, as defined in [Format Compatibility Classes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes)
- [](#VUID-VkTensorViewCreateInfoARM-flags-09745)VUID-VkTensorViewCreateInfoARM-flags-09745  
  If `flags` includes `VK_TENSOR_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM`, the [`descriptorBufferCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-descriptorBufferCaptureReplay) feature **must** be enabled
- [](#VUID-VkTensorViewCreateInfoARM-pNext-09746)VUID-VkTensorViewCreateInfoARM-pNext-09746  
  If the `pNext` chain includes a [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) structure, `flags` **must** contain `VK_TENSOR_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_ARM`
- [](#VUID-VkTensorViewCreateInfoARM-usage-09747)VUID-VkTensorViewCreateInfoARM-usage-09747  
  The `usage` flags of `tensor` **must** have at least one of the following bits set:
  
  - `VK_TENSOR_USAGE_SHADER_BIT_ARM`
  - `VK_TENSOR_USAGE_DATA_GRAPH_BIT_ARM`
- [](#VUID-VkTensorViewCreateInfoARM-usage-09748)VUID-VkTensorViewCreateInfoARM-usage-09748  
  The tensor view’s [format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-tensor-view-format-features) **must** contain the format feature flags required by the `usage` flags of `tensor` for `format` as indicated in the [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#format-feature-dependent-usage-flags](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#format-feature-dependent-usage-flags) section
- [](#VUID-VkTensorViewCreateInfoARM-tensor-09749)VUID-VkTensorViewCreateInfoARM-tensor-09749  
  If `tensor` is non-sparse then it **must** be bound completely and contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object

Valid Usage (Implicit)

- [](#VUID-VkTensorViewCreateInfoARM-sType-sType)VUID-VkTensorViewCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_VIEW_CREATE_INFO_ARM`
- [](#VUID-VkTensorViewCreateInfoARM-pNext-pNext)VUID-VkTensorViewCreateInfoARM-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
- [](#VUID-VkTensorViewCreateInfoARM-sType-unique)VUID-VkTensorViewCreateInfoARM-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkTensorViewCreateInfoARM-flags-parameter)VUID-VkTensorViewCreateInfoARM-flags-parameter  
  `flags` **must** be a valid combination of [VkTensorViewCreateFlagBitsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateFlagBitsARM.html) values
- [](#VUID-VkTensorViewCreateInfoARM-tensor-parameter)VUID-VkTensorViewCreateInfoARM-tensor-parameter  
  `tensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle
- [](#VUID-VkTensorViewCreateInfoARM-format-parameter)VUID-VkTensorViewCreateInfoARM-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [VkTensorViewCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateFlagsARM.html), [vkCreateTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateTensorViewARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorViewCreateInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0