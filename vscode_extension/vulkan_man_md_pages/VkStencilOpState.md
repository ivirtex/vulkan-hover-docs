# VkStencilOpState(3) Manual Page

## Name

VkStencilOpState - Structure specifying stencil operation state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkStencilOpState` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkStencilOpState {
    VkStencilOp    failOp;
    VkStencilOp    passOp;
    VkStencilOp    depthFailOp;
    VkCompareOp    compareOp;
    uint32_t       compareMask;
    uint32_t       writeMask;
    uint32_t       reference;
} VkStencilOpState;
```

## <a href="#_members" class="anchor"></a>Members

- `failOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value specifying the
  action performed on samples that fail the stencil test.

- `passOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value specifying the
  action performed on samples that pass both the depth and stencil
  tests.

- `depthFailOp` is a [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value specifying
  the action performed on samples that pass the stencil test and fail
  the depth test.

- `compareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html) value specifying the
  comparison operator used in the stencil test.

- `compareMask` selects the bits of the unsigned integer stencil values
  participating in the stencil test.

- `writeMask` selects the bits of the unsigned integer stencil values
  updated by the stencil test in the stencil framebuffer attachment.

- `reference` is an integer stencil reference value that is used in the
  unsigned stencil comparison.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkStencilOpState-failOp-parameter"
  id="VUID-VkStencilOpState-failOp-parameter"></a>
  VUID-VkStencilOpState-failOp-parameter  
  `failOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value

- <a href="#VUID-VkStencilOpState-passOp-parameter"
  id="VUID-VkStencilOpState-passOp-parameter"></a>
  VUID-VkStencilOpState-passOp-parameter  
  `passOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html) value

- <a href="#VUID-VkStencilOpState-depthFailOp-parameter"
  id="VUID-VkStencilOpState-depthFailOp-parameter"></a>
  VUID-VkStencilOpState-depthFailOp-parameter  
  `depthFailOp` **must** be a valid [VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html)
  value

- <a href="#VUID-VkStencilOpState-compareOp-parameter"
  id="VUID-VkStencilOpState-compareOp-parameter"></a>
  VUID-VkStencilOpState-compareOp-parameter  
  `compareOp` **must** be a valid [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html),
[VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html),
[VkStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOp.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkStencilOpState"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
