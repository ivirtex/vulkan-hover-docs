# FullyCoveredEXT(3) Manual Page

## Name

FullyCoveredEXT - Indication of whether a fragment is fully covered



## <a href="#_description" class="anchor"></a>Description

`FullyCoveredEXT`  
Decorating a variable with the `FullyCoveredEXT` built-in decoration
will make that variable indicate whether the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-fragment-area"
target="_blank" rel="noopener">fragment area</a> is fully covered by the
generating primitive. This variable is non-zero if conservative
rasterization is enabled and the current fragment area is fully covered
by the generating primitive, and is zero if the fragment is not covered
or partially covered, or conservative rasterization is disabled.

Valid Usage

- <a href="#VUID-FullyCoveredEXT-FullyCoveredEXT-04232"
  id="VUID-FullyCoveredEXT-FullyCoveredEXT-04232"></a>
  VUID-FullyCoveredEXT-FullyCoveredEXT-04232  
  The `FullyCoveredEXT` decoration **must** be used only within the
  `Fragment` `Execution` `Model`

- <a href="#VUID-FullyCoveredEXT-FullyCoveredEXT-04233"
  id="VUID-FullyCoveredEXT-FullyCoveredEXT-04233"></a>
  VUID-FullyCoveredEXT-FullyCoveredEXT-04233  
  The variable decorated with `FullyCoveredEXT` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-FullyCoveredEXT-FullyCoveredEXT-04234"
  id="VUID-FullyCoveredEXT-FullyCoveredEXT-04234"></a>
  VUID-FullyCoveredEXT-FullyCoveredEXT-04234  
  The variable decorated with `FullyCoveredEXT` **must** be declared as
  a boolean value

- <a
  href="#VUID-FullyCoveredEXT-conservativeRasterizationPostDepthCoverage-04235"
  id="VUID-FullyCoveredEXT-conservativeRasterizationPostDepthCoverage-04235"></a>
  VUID-FullyCoveredEXT-conservativeRasterizationPostDepthCoverage-04235  
  If
  `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`conservativeRasterizationPostDepthCoverage`
  is not supported the `PostDepthCoverage` `Execution` `Mode` **must**
  not be declared, when a variable with the `FullyCoveredEXT` decoration
  is declared

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#FullyCoveredEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
