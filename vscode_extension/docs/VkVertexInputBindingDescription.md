# VkVertexInputBindingDescription(3) Manual Page

## Name

VkVertexInputBindingDescription - Structure specifying vertex input binding description



## [](#_c_specification)C Specification

Each vertex input binding is specified by the `VkVertexInputBindingDescription` structure, defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkVertexInputBindingDescription {
    uint32_t             binding;
    uint32_t             stride;
    VkVertexInputRate    inputRate;
} VkVertexInputBindingDescription;
```

## [](#_members)Members

- `binding` is the binding number that this structure describes.
- `stride` is the byte stride between consecutive elements within the buffer.
- `inputRate` is a [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputRate.html) value specifying whether vertex attribute addressing is a function of the vertex index or of the instance index.

## [](#_description)Description

Valid Usage

- [](#VUID-VkVertexInputBindingDescription-binding-00618)VUID-VkVertexInputBindingDescription-binding-00618  
  `binding` **must** be less than `VkPhysicalDeviceLimits`::`maxVertexInputBindings`
- [](#VUID-VkVertexInputBindingDescription-stride-00619)VUID-VkVertexInputBindingDescription-stride-00619  
  `stride` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindingStride`
- [](#VUID-VkVertexInputBindingDescription-stride-04456)VUID-VkVertexInputBindingDescription-stride-04456  
  If the `VK_KHR_portability_subset` extension is enabled, `stride` **must** be a multiple of, and at least as large as, [VkPhysicalDevicePortabilitySubsetPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetPropertiesKHR.html)::`minVertexInputBindingStrideAlignment`

Valid Usage (Implicit)

- [](#VUID-VkVertexInputBindingDescription-inputRate-parameter)VUID-VkVertexInputBindingDescription-inputRate-parameter  
  `inputRate` **must** be a valid [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputRate.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html), [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputRate.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVertexInputBindingDescription)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0