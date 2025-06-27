# VkBlockMatchWindowCompareModeQCOM(3) Manual Page

## Name

VkBlockMatchWindowCompareModeQCOM - Block match window compare modes



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlockMatchWindowCompareModeQCOM.html)
enum describes how block match values within the window are compared.
[VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlockMatchWindowCompareModeQCOM.html)
is defined as:

``` c
// Provided by VK_QCOM_image_processing2
typedef enum VkBlockMatchWindowCompareModeQCOM {
    VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MIN_QCOM = 0,
    VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MAX_QCOM = 1,
} VkBlockMatchWindowCompareModeQCOM;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MIN_QCOM` specifies that windowed
  block match operations return the minimum error within the window.

- `VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MAX_QCOM` specifies that windowed
  block match operations return the maximum error within the window.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QCOM_image_processing2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_image_processing2.html),
[VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBlockMatchWindowCompareModeQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
