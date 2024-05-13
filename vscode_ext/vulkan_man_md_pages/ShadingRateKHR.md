# ShadingRateKHR(3) Manual Page

## Name

ShadingRateKHR - Shading rate of a fragment



## <a href="#_description" class="anchor"></a>Description

`ShadingRateKHR`  
Decorating a variable with the `ShadingRateKHR` built-in decoration will
make that variable contain the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate"
target="_blank" rel="noopener">fragment shading rate</a> for the current
fragment invocation.

Valid Usage

- <a href="#VUID-ShadingRateKHR-ShadingRateKHR-04490"
  id="VUID-ShadingRateKHR-ShadingRateKHR-04490"></a>
  VUID-ShadingRateKHR-ShadingRateKHR-04490  
  The `ShadingRateKHR` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-ShadingRateKHR-ShadingRateKHR-04491"
  id="VUID-ShadingRateKHR-ShadingRateKHR-04491"></a>
  VUID-ShadingRateKHR-ShadingRateKHR-04491  
  The variable decorated with `ShadingRateKHR` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-ShadingRateKHR-ShadingRateKHR-04492"
  id="VUID-ShadingRateKHR-ShadingRateKHR-04492"></a>
  VUID-ShadingRateKHR-ShadingRateKHR-04492  
  The variable decorated with `ShadingRateKHR` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ShadingRateKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
