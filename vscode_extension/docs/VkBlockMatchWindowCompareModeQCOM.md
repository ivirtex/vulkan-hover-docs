# VkBlockMatchWindowCompareModeQCOM(3) Manual Page

## Name

VkBlockMatchWindowCompareModeQCOM - Block match window compare modes



## [](#_c_specification)C Specification

The [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlockMatchWindowCompareModeQCOM.html) enum describes how block match values within the window are compared. [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlockMatchWindowCompareModeQCOM.html) is defined as:

```c++
// Provided by VK_QCOM_image_processing2
typedef enum VkBlockMatchWindowCompareModeQCOM {
    VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MIN_QCOM = 0,
    VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MAX_QCOM = 1,
} VkBlockMatchWindowCompareModeQCOM;
```

## [](#_description)Description

- `VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MIN_QCOM` specifies that windowed block match operations return the minimum error within the window.
- `VK_BLOCK_MATCH_WINDOW_COMPARE_MODE_MAX_QCOM` specifies that windowed block match operations return the maximum error within the window.

## [](#_see_also)See Also

[VK\_QCOM\_image\_processing2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_image_processing2.html), [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBlockMatchWindowCompareModeQCOM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0