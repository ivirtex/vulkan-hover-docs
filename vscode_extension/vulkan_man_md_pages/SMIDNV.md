# SMIDNV(3) Manual Page

## Name

SMIDNV - SM ID on which a shader invocation is running



## <a href="#_description" class="anchor"></a>Description

`SMIDNV`  
Decorating a variable with the `SMIDNV` built-in decoration will make
that variable contain the ID of the SM on which the current shader
invocation is running. This variable is in the range \[0,
`SMCountNV`-1\].

Valid Usage

- <a href="#VUID-SMIDNV-SMIDNV-04365" id="VUID-SMIDNV-SMIDNV-04365"></a>
  VUID-SMIDNV-SMIDNV-04365  
  The variable decorated with `SMIDNV` **must** be declared using the
  `Input` `Storage` `Class`

- <a href="#VUID-SMIDNV-SMIDNV-04366" id="VUID-SMIDNV-SMIDNV-04366"></a>
  VUID-SMIDNV-SMIDNV-04366  
  The variable decorated with `SMIDNV` **must** be declared as a scalar
  32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SMIDNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
