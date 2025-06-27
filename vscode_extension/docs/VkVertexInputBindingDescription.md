# VkVertexInputBindingDescription(3) Manual Page

## Name

VkVertexInputBindingDescription - Structure specifying vertex input
binding description



## <a href="#_c_specification" class="anchor"></a>C Specification

Each vertex input binding is specified by the
`VkVertexInputBindingDescription` structure, defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkVertexInputBindingDescription {
    uint32_t             binding;
    uint32_t             stride;
    VkVertexInputRate    inputRate;
} VkVertexInputBindingDescription;
```

## <a href="#_members" class="anchor"></a>Members

- `binding` is the binding number that this structure describes.

- `stride` is the byte stride between consecutive elements within the
  buffer.

- `inputRate` is a [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html) value
  specifying whether vertex attribute addressing is a function of the
  vertex index or of the instance index.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkVertexInputBindingDescription-binding-00618"
  id="VUID-VkVertexInputBindingDescription-binding-00618"></a>
  VUID-VkVertexInputBindingDescription-binding-00618  
  `binding` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-VkVertexInputBindingDescription-stride-00619"
  id="VUID-VkVertexInputBindingDescription-stride-00619"></a>
  VUID-VkVertexInputBindingDescription-stride-00619  
  `stride` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxVertexInputBindingStride`

- <a href="#VUID-VkVertexInputBindingDescription-stride-04456"
  id="VUID-VkVertexInputBindingDescription-stride-04456"></a>
  VUID-VkVertexInputBindingDescription-stride-04456  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, `stride` **must** be a multiple of, and at least
  as large as,
  [VkPhysicalDevicePortabilitySubsetPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetPropertiesKHR.html)::`minVertexInputBindingStrideAlignment`

Valid Usage (Implicit)

- <a href="#VUID-VkVertexInputBindingDescription-inputRate-parameter"
  id="VUID-VkVertexInputBindingDescription-inputRate-parameter"></a>
  VUID-VkVertexInputBindingDescription-inputRate-parameter  
  `inputRate` **must** be a valid
  [VkVertexInputRate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html),
[VkVertexInputRate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputRate.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVertexInputBindingDescription"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
