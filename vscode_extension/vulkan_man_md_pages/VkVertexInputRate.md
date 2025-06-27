# VkVertexInputRate(3) Manual Page

## Name

VkVertexInputRate - Specify rate at which vertex attributes are pulled
from buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription.html)::`inputRate`,
specifying the rate at which vertex attributes are pulled from buffers,
are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkVertexInputRate {
    VK_VERTEX_INPUT_RATE_VERTEX = 0,
    VK_VERTEX_INPUT_RATE_INSTANCE = 1,
} VkVertexInputRate;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VERTEX_INPUT_RATE_VERTEX` specifies that vertex attribute
  addressing is a function of the vertex index.

- `VK_VERTEX_INPUT_RATE_INSTANCE` specifies that vertex attribute
  addressing is a function of the instance index.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription.html),
[VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVertexInputRate"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
